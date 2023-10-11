package sdk

type Response struct {
	Code      int    `json:"code,omitempty"`
	Success   string `json:"success,omitempty"`
	Message   string `json:"message,omitempty"`
	MessageId int    `json:"messageId,omitempty"`
}
