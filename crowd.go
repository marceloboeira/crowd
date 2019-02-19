package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julienschmidt/httprouter"
)

var port int
var endpoint string
var queue string

func init() {
	flag.IntVar(&port, "port", 2104, "HTTP Server Port")
	flag.StringVar(&endpoint, "endpoint", "/api/foo", "HTTP path to receive POST requests")
	flag.StringVar(&queue, "queue", "https://sqs.eu-central-1.amazonaws.com/21042018/foo", "URL for the SQS queue")
	flag.Parse()
}

func main() {
	q := sqs.New(session.Must(session.NewSession(&aws.Config{})))

	router := httprouter.New()
	router.POST(endpoint, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		_, err = q.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String(string(body)),
			QueueUrl:    aws.String(queue),
		})

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})

	address := fmt.Sprintf(":%d", port)
	log.Println("crowd is running at", address, "with", endpoint, "-->", queue)

	http.ListenAndServe(address, router)
}
