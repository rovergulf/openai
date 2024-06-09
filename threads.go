package openai

const (
	threadUrl    = "/v1/threads"
	threadUrlFmt = "/v1/threads/%s"
)

type Thread struct {
	Id       string     `json:"id"`
	Type     string     `json:"type"`
	Messages []*Message `json:"messages"`
}
