package sink

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type mockSQS struct {
	sqsiface.SQSAPI
	Counter int
	Resp    sqs.SendMessageOutput
}

func (m *mockSQS) SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	m.Counter++

	return &m.Resp, nil
}

func TestSQS(t *testing.T) {
	mock := mockSQS{Counter: 0}

	q := &SQS{
		Client: &mock,
		Url:    "http://localhost:9324/sqs/123",
	}

	q.Push([]byte("test"))
	q.Push([]byte("test"))
	q.Push([]byte("test"))

	if mock.Counter != 3 {
		t.Errorf("SQS did not receive push")
	}
}
