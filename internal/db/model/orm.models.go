package model

func Tables() (t []interface{}) {
	t = append(t,
		new(MessageInfo),
		new(MessageTask),
		new(MessageRunningTask),
		new(OpenClient),
		new(SmsMessageInfo),
		new(ApnsMessageInfo),
	)
	return
}
