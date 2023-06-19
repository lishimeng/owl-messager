package ms

import "errors"

/**
发送的数据格式：
{
  "message": {
    "subject": "This is Te_st email sent by Graph API",
    "body": {
      "contentType": "Text",
      "content": "Do not get over-excited. - Tim BAO"
    },
    "toRecipients": [
      {
        "emailAddress": {
          "address": "tim.bao@basf.com"
        }
      },
      {
        "emailAddress": {
          "address": "fiona.fang@basf.com"
        }
	  },
      {
        "emailAddress": {
          "address": "mingxing-star.chen@basf-ypc.com.cn"
        }
	  }
    ]
  },
  "saveToSentItems": "true"
}

*/

type MessageBody struct {
	ContentType string `json:"contentType,omitempty"`
	Content     string `json:"content,omitempty"`
}

type Addr struct {
	Address string `json:"address"`
}

type MessageReceiver struct {
	EmailAddress Addr `json:"emailAddress"`
}

type Message struct {
	Subject      string            `json:"subject"`
	Body         MessageBody       `json:"body"`
	ToRecipients []MessageReceiver `json:"toRecipients"`
}

type MessageWrapper struct {
	Message         Message `json:"message"`
	SaveToSentItems string  `json:"saveToSentItems"`
}

func BuildMessage(subject string, body string, receivers ...string) (m MessageWrapper, err error) {

	if len(receivers) <= 0 {
		err = errors.New("no receiver")
		return
	}
	if len(subject) == 0 {
		err = errors.New("no subject")
		return
	}
	if len(body) == 0 {
		err = errors.New("body nil")
		return
	}
	m.SaveToSentItems = "false"
	m.Message = Message{
		Subject: subject,
		Body: MessageBody{
			ContentType: "HTML", // 枚举：Text/HTML
			Content:     body,
		},
	}

	var rs []MessageReceiver
	for _, r := range receivers {
		rs = append(rs, MessageReceiver{
			EmailAddress: Addr{
				Address: r,
			},
		})
	}
	m.Message.ToRecipients = rs
	return
}
