package ingestion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rphmauriciodev/myFinance/internal/dtos"
)

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body dtos.Transaction

	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"error":"invalid body"}`,
		}, nil
	}

	if err := body.Validate(); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf(`{"error":"%s"}`, err.Error()),
		}, nil
	}

	if err := publishTransaction(body); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"error":"failed to publish transaction"}`,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       `{"message":"OK"}`,
	}, nil
}

func publishTransaction(transaction dtos.Transaction) error {
	// Implement the logic to publish the transaction here
	// For example, you might want to send it to a message queue,
	// notify other services, or perform other actions.

	return nil
}
