package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	fmt.Println("This message will show up in the CLI console.")

	client := &http.Client{}
	// If I needed a secret API key here, I could puck it out of an environment variable.
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyText[:])
	// fmt.Printf("%s\n", bodyText)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/json"},
		Body:            bodyString,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
