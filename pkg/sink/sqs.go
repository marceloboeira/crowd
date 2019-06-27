package sink

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type SQS struct {
	Client sqsiface.SQSAPI
	URL    string
}

func NewSQS(url string) SQS {
	queue := sqs.New(session.Must(session.NewSession(&aws.Config{})))

	return SQS{Client: queue, URL: url}
}

func (s SQS) Push(payload []byte) error {
	_, err := s.Client.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(payload)),
		QueueUrl:    aws.String(s.URL),
	})

	return err
}
