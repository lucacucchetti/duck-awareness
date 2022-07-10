package main

import (
	"flag"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "You are on the Duck Awareness path now...")
	})
	localRun := flag.Bool("localRun", false, "whether it is a local run and should use test env")
	flag.Parse()

	if *localRun == true {
		http.ListenAndServe(":80", http.DefaultServeMux)
	} else {
		lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
	}
}
