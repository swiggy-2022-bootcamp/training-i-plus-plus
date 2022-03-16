package main

import (
	//"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"dynamo/user"
	"net/http"
)

// func CreateLocalClient() *dynamodb.Client {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("mumbai"),
// 		config.WithEndpointResolver(aws.EndpointResolverFunc(
// 			func(service, region string) (aws.Endpoint, error) {
// 				return aws.Endpoint{URL: "http://localhost:8042"}, nil
// 			})),
// 		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
// 			Value: aws.Credentials{
// 				AccessKeyID: "fake", SecretAccessKey: "fake",
// 			},
// 		}),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return dynamodb.NewFromConfig(cfg)
// }

func main() {
	s := user.NewServer()
	// id := uuid.New().String()
	// u2 := user.User{
	// 	UUID:  id,
	// 	Name:  "suhas ravishanker",
	// 	Email: "suhasravi@google.io",
	// }
	// user.AddUser(S.Client, u2)
	http.ListenAndServe(":8081", s)
}
