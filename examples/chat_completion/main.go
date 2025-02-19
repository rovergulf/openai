package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/rovergulf/openai"
)

// https://platform.openai.com/docs/api-reference/chat/create
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(apiKey)

	req := &openai.CompletionRequest{
		Model: openai.GPT4o,
		Messages: []*openai.CompletionMessage{
			{
				Role:    openai.SystemRole,
				Content: openai.DefaultSystemTask,
			},
			{
				Role:    openai.UserRole,
				Content: "Hello there! Who was the father of Luke Skywalker?",
			},
		},
	}

	completionResult, err := client.ChatCompletion(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	contents := completionResult.Choices[0].Message.Content

	fmt.Println(contents)
}
