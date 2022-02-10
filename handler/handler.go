package handler

func New() *WebhookHandler {
	return &WebhookHandler{}
}

type WebhookHandler struct {
	webhooks map[string]string
}

func (h *WebhookHandler) Handle(path string, body []byte) error {
	return nil
}
