package main

import (
	"example/hello/src/entities"
	"example/hello/src/presentation"
	"flag"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	duck_awareness_phrase := "You are on the Duck Awareness path now..."
	testing_board := entities.NewBoard(duck_awareness_phrase, duck_awareness_phrase, []*entities.Resource{})
	testing_board.AddResource(testing_board)
	var presenter presentation.Presenter = presentation.TextPresenter{}
	_, board_representation, _ := presenter.DisplayBoard(*testing_board)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, board_representation)
	})
	localRun := flag.Bool("localRun", false, "whether it is a local run and should use test env")
	flag.Parse()

	// runs with "go run main.go --localRun" from the directory of main.go
	if *localRun == true {
		http.ListenAndServe(":80", http.DefaultServeMux)
	} else {
		lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
	}
}
