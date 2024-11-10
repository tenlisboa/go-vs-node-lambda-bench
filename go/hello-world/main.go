package main

import (
	"fmt"
	"math"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Função para verificar se um número é primo
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Função para calcular números primos até um limite
func calculatePrimes(limit int) []int {
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	limit := 1000000
	primes := calculatePrimes(limit)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Números primos até %d: %d\n", limit, len(primes)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
