package main

import (
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "You are on the Duck Awareness path now...")
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}