package handlers

import "github.com/aws/aws-lambda-go/events"

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error){
	resp := events.APIGatewayProxyResponse{Headers: map[string]string[]}
}
