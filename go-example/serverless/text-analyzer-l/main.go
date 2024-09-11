package main

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Source        string `json:"source"`
	WordCount     int    `json:"word_count"`
	CharCount     int    `json:"char_count"`
	SentenceCount int    `json:"sentence_count"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request: unable to parse JSON",
		}, nil
	}

	if req.Text == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request: missing text",
		}, nil
	}

	words := strings.Fields(req.Text)
	wordCount := len(words)
	charCount := len(req.Text)

	re := regexp.MustCompile(`\w+[.!?]`)
	sentenceCount := len(re.FindAllString(req.Text, -1))

	response := Response{
		WordCount:     wordCount,
		CharCount:     charCount,
		SentenceCount: sentenceCount,
		Source:        "serverless",
	}

	jsonResponse, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
