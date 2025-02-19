package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/rovergulf/openai"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apiKey := os.Getenv("OPENAI_API_KEY")

	var assistantId, threadId string

	flag.StringVar(&assistantId, "assistant-id", "", "")
	flag.StringVar(&threadId, "thread-id", "", "")
	flag.Parse()

	client := openai.NewClient(apiKey)

	assistant, err := client.RetrieveAssistant(ctx, assistantId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(assistant)

	thread, err := client.RetrieveThread(ctx, threadId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(thread)
}
