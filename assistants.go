package openai

const (
	assistantsUrl = "/v1/assistants"
	assistantUrl  = "/v1/assistants/%s"
)

type Assistant struct {
	Id             string  `json:"id"`
	Object         string  `json:"object"`
	Model          string  `json:"model"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Instructions   string  `json:"instructions"`
	Tools          []Tool  `json:"tools"`
	Temperature    float32 `json:"temperature"`
	TopP           float32 `json:"top_p"`
	ResponseFormat string  `json:"response_format"`
}
