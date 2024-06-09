package openai

type ModelResponse struct {
	Object string   `json:"object"`
	Data   []*Model `json:"data"`
}

type Model struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int64  `json:"created"`
	OwnedBy   string `json:"owned_by"`
}
