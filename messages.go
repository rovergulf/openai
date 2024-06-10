package openai

const (
	messagesUrlFmt = "/v1/threads/%s/messages"
	messageUrlFmt  = "/v1/threads/%s/messages/%s"
)

type MessagesRequest struct {
	ListRequestParams
	RunId string `json:"run_id"`
}

type MessageAttachment struct {
	FileId string `json:"file_id"`
	Tools  []Tool `json:"tools,omitempty"`
}

type MessageContent struct {
	Type string      `json:"type"`
	Text MessageText `json:"text"`
}

type MessageText struct {
	Value       string `json:"value"`
	Annotations []any  `json:"annotations"`
}

type Message struct {
	Id          string           `json:"id"`
	Object      string           `json:"object"`
	CreatedAt   int              `json:"created_at"`
	AssistantId string           `json:"assistant_id"`
	ThreadId    string           `json:"thread_id"`
	RunId       string           `json:"run_id"`
	Role        string           `json:"role"`
	Content     []MessageContent `json:"content"`
}
