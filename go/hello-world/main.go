package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func encryptAES(key, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return hex.EncodeToString(ciphertext), nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := []byte("thisis32bitlongpassphraseimusing")
	plaintext := []byte("Hello, AWS Lambda!")
	encryptedResults := make([]string, 0)

	start := time.Now()

	for i := 0; i < 10000; i++ {
		encrypted, err := encryptAES(key, plaintext)
		if err != nil {
			log.Fatal("Erro na criptografia:", err)
		}
		encryptedResults = append(encryptedResults, encrypted)
	}

	duration := time.Since(start)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("O tempo de execução foi de: %s, resultados criptografados: %v", duration, len(encryptedResults)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
