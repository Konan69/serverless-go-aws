package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is required")
	}

	return fmt.Sprintf("succesfully processed - %s", event.Username) , nil 
}



func main() {
	_ := app.NewApp()
	lambda.Start(HandleRequest)

} 