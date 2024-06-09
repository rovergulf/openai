package openai

const (
	chatCompletionUrl = "/v1/chat/completions"
)

type CompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Model       string               `json:"model"`
	Messages    []*CompletionMessage `json:"messages"`
	Seed        int64                `json:"seed,omitempty"`
	Temperature float32              `json:"temperature,omitempty"`
}

type CompletionResponse struct {
	Id                string    `json:"id"`
	Object            string    `json:"object"`
	CreatedAt         int64     `json:"created"`
	Model             string    `json:"model"`
	Choices           []*Choice `json:"choices"`
	Usage             *Usage    `json:"usage"`
	SystemFingerprint string    `json:"system_fingerprint"`
}

type Choice struct {
	Index   int                `json:"index"`
	Message *CompletionMessage `json:"message"`
	Finish  string             `json:"finish_reason"`
}
