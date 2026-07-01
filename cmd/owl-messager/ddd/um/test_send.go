package um

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lishimeng/owl-messager/internal/db/repo"
	messagerSender "github.com/lishimeng/owl-messager/internal/messager/sender"
)

func marshalTestParams(params any) (string, error) {
	if params == nil {
		return "{}", nil
	}
	switch v := params.(type) {
	case string:
		if v == "" {
			return "{}", nil
		}
		return v, nil
	default:
		bs, err := json.Marshal(params)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}
}

// SendMailNow creates a mail message and sends it synchronously (console test send).
func SendMailNow(tenantCode, templateCode, subject, receiver string, params any) (messageId int, err error) {
	paramsJSON, err := marshalTestParams(params)
	if err != nil {
		return
	}
	if subject == "" {
		subject = DefaultTitle
	}
	m, err := serviceAddMail(tenantCode, templateCode, paramsJSON, subject, receiver, nil)
	if err != nil {
		return
	}
	mailInfo, err := repo.GetMailByMessageId(m.Id)
	if err != nil {
		return
	}
	s, err := messagerSender.NewMailSender(context.Background())
	if err != nil {
		return
	}
	err = s.Send(mailInfo)
	messageId = m.Id
	return
}

// SendSmsNow creates an SMS message and sends it synchronously.
func SendSmsNow(tenantCode, templateCode, receiver string, params any) (messageId int, err error) {
	paramsJSON, err := marshalTestParams(params)
	if err != nil {
		return
	}
	m, err := serviceAddSms(tenantCode, templateCode, paramsJSON, receiver)
	if err != nil {
		return
	}
	smsInfo, err := repo.GetSmsByMessageId(m.Id)
	if err != nil {
		return
	}
	s, err := messagerSender.NewSmsSender(context.Background())
	if err != nil {
		return
	}
	err = s.Send(smsInfo)
	messageId = m.Id
	return
}

// SendImNow creates an IM message and sends it synchronously.
func SendImNow(tenantCode, templateCode, receiver string, params any) (messageId int, err error) {
	paramsJSON, err := marshalTestParams(params)
	if err != nil {
		return
	}
	m, err := serviceAddIm(tenantCode, templateCode, paramsJSON, receiver)
	if err != nil {
		return
	}
	imInfo, err := repo.GetImByMessageId(m.Id)
	if err != nil {
		return
	}
	s, err := messagerSender.NewImSender(context.Background())
	if err != nil {
		return
	}
	err = s.Send(imInfo)
	messageId = m.Id
	return
}

func FormatSendError(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%v", err)
}
