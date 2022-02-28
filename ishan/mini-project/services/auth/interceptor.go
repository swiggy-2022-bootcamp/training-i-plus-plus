package auth

import (
	"context"
	"fmt"
	"log"
	jwtmanager "swiggy/train_reservation/helpers/lib"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	JWTManager      *jwtmanager.JWTManager
	AccessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *jwtmanager.JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

func AccessibleRoles() map[string][]string {
	const servicePath = "/auth.AuthService/"

	return map[string][]string{
		servicePath + "CheckAuth": {"admin", "user"},
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.AccessibleRoles[method]
	if !ok {
		// everyone can access
		fmt.Println("Escape Auth Method")
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.JWTManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
