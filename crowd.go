package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	q := sqs.New(session.New())
	http.HandleFunc("/api/leads", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		_, err = q.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String(string(body)),
			QueueUrl:    aws.String("https://sqs.eu-central-1.amazonaws.com/046001896437/test"),
		})

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})

	http.ListenAndServe(":9000", nil)
}
