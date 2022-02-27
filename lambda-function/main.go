package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var ginLambda *ginadapter.GinLambda
var router *gin.Engine

// init the Gin Server
func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router = gin.Default()
	router.GET("/ping", ping)
	router.GET("/pong", pong)
	router.GET("/calculator", calculator)

	ginLambda = ginadapter.New(router)
}

// Handler will deal with Gin working with Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" {
		// Run on AWS Lambda
		lambda.Start(Handler)
	} else {
		// Run locally
		router.Run()
	}
}
