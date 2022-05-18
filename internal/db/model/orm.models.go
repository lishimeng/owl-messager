package model

func Tables() (t []interface{}) {
	t = append(t,
		new(MessageInfo),
		new(MessageTask),
		new(MessageRunningTask),

		new(OpenClient),

		new(MailMessageInfo),
		new(MailTemplateInfo),
		new(MailSenderInfo),

		new(SmsMessageInfo),
		new(SmsSenderInfo),
		new(SmsTemplateInfo),

		new(ApnsMessageInfo),
		new(ApnsSenderInfo),
	)
	return
}
