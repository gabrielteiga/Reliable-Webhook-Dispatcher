package webhooks

type WebhooksRequest struct {
	TargetUrl string                 `json:"target_url"`
	Event     string                 `json:"event"`
	Payload   WebhooksRequestPayload `json:"payload"`
}

type WebhooksRequestPayload struct {
	PaymentID string `json:"payment_id"`
	Amount    int    `json:"amount"`
}
