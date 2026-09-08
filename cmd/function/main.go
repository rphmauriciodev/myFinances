package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rphmauriciodev/myFinance/internal/ingestion"
)

func main() {

	fmt.Println("Starting the application...")

	lambda.Start(ingestion.Handler)
}
