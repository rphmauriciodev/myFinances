package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Queue struct {
	client   *sqs.Client
	queueURL string
}

func NewQueue(ctx context.Context, queueURL string) (*Queue, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return &Queue{
		client:   client,
		queueURL: queueURL,
	}, nil
}

func (q *Queue) Send(ctx context.Context, body string) error {
	_, err := q.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(q.queueURL),
		MessageBody: aws.String(body),
	})

	return err
}
