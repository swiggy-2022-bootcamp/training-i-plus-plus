package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {

	type Movie struct {
		Year  int         `json:"year"`
		Title string      `json:"title"`
		Info  interface{} `json:"info"`
		UUID  int         `json:"id"`
	}

	moviesData, err := os.Open("moviedata.json")
	defer moviesData.Close()
	if err != nil {
		fmt.Println("Could not open the moviedata.json file", err.Error())
		os.Exit(1)
	}

	var movies []Movie
	err = json.NewDecoder(moviesData).Decode(&movies)
	if err != nil {
		fmt.Println("Could not decode the moviedata.json data", err.Error())
		os.Exit(1)
	}

	config := &aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("dynamodb.us-west-2.amazonaws.com"),
	}

	sess := session.Must(session.NewSession(config))

	svc := dynamodb.New(sess)

	for _, movie := range movies {

		info, err := dynamodbattribute.MarshalMap(movie)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal the movie, %v", err))
		}

		input := &dynamodb.PutItemInput{
			Item:      info,
			TableName: aws.String("Movies"),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	}

	fmt.Printf("We have processed %v records\n", len(movies))
}
