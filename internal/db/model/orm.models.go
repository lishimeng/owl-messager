package model

func Tables() (t []interface{}) {
	t = append(t,
		new(MessageInfo),
		new(MailMessageInfo),
		new(SmsMessageInfo),
		new(SmsSenderInfo),
		new(MailSenderInfo),
		new(MailTemplateInfo),
		new(MessageTask),
		new(MessageRunningTask),
	)
	return
}
