package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rphmauriciodev/myFinance/internal/ingestion"
	"github.com/rphmauriciodev/myFinance/internal/queue"
	"github.com/spf13/viper"
)

func main() {

	ctx := context.Background()

	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			slog.Error("erro lendo arquivo de configuração: %w", err)
			return
		}
	}

	slog.Info("Starting the application...")

	var queueURL = v.GetString("SQS_QUEUE_URL")

	queue, err := queue.NewQueue(ctx, queueURL)
	if err != nil {
		slog.Error("Failed to create SQS queue: %w", err)
		return
	}

	handler := ingestion.NewHandler(queue)

	slog.Info("Lambda initialized successfully")

	lambda.Start(handler.Handle)

}
