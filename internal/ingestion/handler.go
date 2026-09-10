package ingestion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rphmauriciodev/myFinance/internal/dtos"
	"github.com/rphmauriciodev/myFinance/internal/queue"
)

type Handler struct {
	queue *queue.Queue
}

func NewHandler(queue *queue.Queue) *Handler {
	return &Handler{
		queue: queue,
	}
}

func (h *Handler) Handle(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
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

	if err := h.queue.Send(ctx, req.Body); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf(`{"error":"failed to publish transaction: %s"}`, err.Error()),
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       `{"message":"OK"}`,
	}, nil
}
