package main

import (
	"context"
	"fmt"
	"log"
	authpb "swiggy/train_reservation/services/auth/authpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authClient  *AuthClient
	authMethods map[string]bool
	accessToken string
}

type AuthClient struct {
	service  authpb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := authpb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

func NewAuthInterceptor(
	authClient *AuthClient,
	authMethods map[string]bool,
	refreshDuration time.Duration,
) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authClient:  authClient,
		authMethods: authMethods,
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken
	log.Printf("token refreshed: %v", accessToken)

	return nil
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err := interceptor.refreshToken()
	if err != nil {
		return err
	}

	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

func authMethods() map[string]bool {
	const servicePath = "/auth.AuthService/"

	return map[string]bool{
		servicePath + "CheckAuth": true,
	}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &authpb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

func signUp(c authpb.AuthServiceClient) {
	req := &authpb.SignupRequest{
		Username: "Ishan",
		Password: "123",
	}
	res, err := c.Signup(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Auth: %v", res.GetId())
}

func testAuth(c authpb.AuthServiceClient) {
	res, err := c.CheckAuth(context.Background(), &authpb.AuthRequest{})
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Auth: %v", res.GetStatus())
}

const (
	username = "ishan"
	password = "123"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc1, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("coul not connet %f", err)
	}
	defer cc1.Close()

	authClient := NewAuthClient(cc1, username, password)
	interceptor, err := NewAuthInterceptor(authClient, authMethods(), 40)

	cc2, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	c := authpb.NewAuthServiceClient(cc2)
	//login(c)
	// signUp(c)
	testAuth(c)
}
