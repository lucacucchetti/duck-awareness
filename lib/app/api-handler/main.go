package main

import (
	"context"
	controllers "example/hello/src/controllers"
	. "example/hello/src/entities"
	"flag"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
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

	var repo Repository = MockRepository{}

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := gin.Default()
	registerEndpoints(router, repo)

	ginLambda = ginadapter.New(router)

	// runs with "go run main.go --localRun" from the directory of main.go
	if *localRun == true {
		router.Run()
	} else {
		lambda.Start(Handler)
	}
}

func registerEndpoints(router *gin.Engine, repo Repository) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are on the Duck Awareness path now...",
		})
	})

	router.GET("/board/:boardId", controllers.GetBoard(repo))
	router.GET("/text_note/:noteId", controllers.GetTextNote(repo))
}
