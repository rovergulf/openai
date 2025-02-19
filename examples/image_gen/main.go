package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/rovergulf/openai"
)

// https://platform.openai.com/docs/api-reference/images/create
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(apiKey)

	req := &openai.ImageGenRequest{
		Model:  openai.Dalle3,
		Prompt: "Draw a monkey with banana and fruits",
		Size:   openai.Size1024x1024,
		Amount: 1,
		Style:  openai.NaturalStyle,
		//ResponseFormat: "base64", // json
	}

	imageGenResponse, err := client.ImageGeneration(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(imageGenResponse.Data[0].Url)
}
