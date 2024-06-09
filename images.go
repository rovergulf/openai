package openai

import "os"

const (
	Dalle2 = "dall-e-2"
	Dalle3 = "dall-e-3"

	Size1024x1024 = "1024x1024"
	Size1792x1024 = "1792x1024"
	Size1024x1792 = "1024x1792"

	VividStyle   = "vivid"
	NaturalStyle = "natural"
)

const (
	imageGenUrl        = "/v1/images/generations"
	imageEditUrl       = "/v1/images/edit"
	imageVariationsUrl = "/v1/images/variations"
	audioSpeechUrl     = "/v1/audio/speech"
)

type ImageGenRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Size   string `json:"size"`
	Amount int64  `json:"n,omitempty"`
	Style  string `json:"style,omitempty"`
	//ResponseFormat string `json:"response_format,omitempty"`
}

type ImageGenResponse struct {
	CreatedAt int64               `json:"created"`
	Data      []*ImageGenDataItem `json:"data"`
}

type ImageGenDataItem struct {
	Url           string `json:"url"`
	RevisedPrompt string `json:"revised_prompt"`
	B64Json       string `json:"b64_json"`
}

type ImageEditRequest struct {
	Image  os.File `json:"image"`
	Mask   os.File `json:"mask"`
	Model  string  `json:"model"`
	Prompt string  `json:"prompt"`
	Amount int64   `json:"n"`
}

type ImageVariationRequest struct{}
