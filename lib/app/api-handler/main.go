package main

import (
	"flag"
	"log"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	localRun := flag.Bool("localRun", false, "whether it is a local run and should use test env")
	flag.Parse()

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are on the Duck Awareness path now...",
		})
	})

	ginLambda = ginadapter.New(router)

	// runs with "go run main.go --localRun" from the directory of main.go
	if *localRun == true {
		router.Run()
	} else {
		lambda.Start(Handler)
	}
}

