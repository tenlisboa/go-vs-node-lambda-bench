package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func simulateServiceLatency(min, max int) time.Duration {
	latency := rand.Intn(max-min+1) + min
	t := time.Duration(latency) * time.Millisecond
	time.Sleep(t)

	return time.Duration(latency)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rand.Seed(time.Now().UnixNano())

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("O serviço demorou %d ms para responder", simulateServiceLatency(600, 1000)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
