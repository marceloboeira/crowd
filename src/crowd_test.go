package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"

	"github.com/julienschmidt/httprouter"
)

type mockSQS struct {
	sqsiface.SQSAPI
	Resp sqs.SendMessageOutput
}

func (m mockSQS) SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return &m.Resp, nil
}

func TestCrowd(t *testing.T) {
	q := Queue{
		Client: mockSQS{},
		URL:    "http://localhost:9324",
	}
	crowd := Crowd{q: q}
	router := httprouter.New()
	router.POST("/api/foo", crowd.Handle)
	data := url.Values{}
	data.Add("foo", "bar")

	req, _ := http.NewRequest("POST", "/api/foo", bytes.NewBufferString(data.Encode()))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned with wrong status. Got %d, expected %d", status, http.StatusOK)
	}
}
