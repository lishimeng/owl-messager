package sdk

type Client interface {
	SendMail(req MailRequest) (resp Response, err error)
	SendSms(req SmsRequest) (resp Response, err error)
	SendApns(req ApnsRequest) (resp Response, err error)
}
