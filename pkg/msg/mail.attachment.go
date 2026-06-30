package msg

import "encoding/json"

// MailAttachmentRef 邮件附件引用（元数据在 mail_attachment 表，文件在 staging 目录）
type MailAttachmentRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Mime string `json:"mime,omitempty"`
	Size int64  `json:"size"`
}

func ParseMailAttachmentRefs(raw string) (refs []MailAttachmentRef, err error) {
	if raw == "" {
		return nil, nil
	}
	err = json.Unmarshal([]byte(raw), &refs)
	return
}

func MarshalMailAttachmentRefs(refs []MailAttachmentRef) string {
	if len(refs) == 0 {
		return ""
	}
	b, _ := json.Marshal(refs)
	return string(b)
}
