package sdk

type Response struct {
	Code      interface{} `json:"code,omitempty"`
	Success   interface{} `json:"success,omitempty"`
	Message   string      `json:"message,omitempty"`
	MessageId int         `json:"messageId,omitempty"`
}
