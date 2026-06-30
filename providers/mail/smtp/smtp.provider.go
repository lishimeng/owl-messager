package smtp

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"mime"
	"net/smtp"
	"strings"

	"github.com/lishimeng/owl-messager/pkg/msg"
)

type AttachmentPart struct {
	Filename    string
	ContentType string
	Data        []byte
}

type MailSmtpProvider struct {
	Config msg.SmtpConfig
}

func encodeFrom(email, alias string) string {
	if alias == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", alias, email)
}

func buildMessage(from, subject, htmlBody string, to []string, attachments []AttachmentPart) []byte {
	if len(attachments) == 0 {
		return buildSimpleMessage(from, subject, htmlBody, to)
	}
	boundary := "owl-messager-" + randomBoundaryToken()
	var buf bytes.Buffer
	buf.WriteString("From: " + from + "\r\n")
	buf.WriteString("To: " + strings.Join(to, ", ") + "\r\n")
	buf.WriteString("Subject: " + subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n")
	buf.WriteString("\r\n")

	writePart(&buf, boundary, "text/html; charset=UTF-8", "", htmlBody)
	for _, att := range attachments {
		ct := att.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}
		filename := mime.QEncoding.Encode("utf-8", att.Filename)
		headers := "Content-Type: " + ct + `; name="` + filename + `"` + "\r\n" +
			"Content-Transfer-Encoding: base64\r\n" +
			`Content-Disposition: attachment; filename="` + filename + `"`
		writePart(&buf, boundary, "", headers, base64.StdEncoding.EncodeToString(att.Data))
	}
	buf.WriteString("--" + boundary + "--\r\n")
	return buf.Bytes()
}

func buildSimpleMessage(from, subject, htmlBody string, to []string) []byte {
	var buf bytes.Buffer
	buf.WriteString("From: " + from + "\r\n")
	buf.WriteString("To: " + strings.Join(to, ", ") + "\r\n")
	buf.WriteString("Subject: " + subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	buf.WriteString("\r\n")
	buf.WriteString(htmlBody)
	return buf.Bytes()
}

func writePart(buf *bytes.Buffer, boundary, contentType, extraHeaders, body string) {
	buf.WriteString("--" + boundary + "\r\n")
	if contentType != "" {
		buf.WriteString("Content-Type: " + contentType + "\r\n")
	}
	if extraHeaders != "" {
		buf.WriteString(extraHeaders + "\r\n")
	}
	buf.WriteString("\r\n")
	buf.WriteString(body)
	buf.WriteString("\r\n")
}

func randomBoundaryToken() string {
	var b [12]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func (s MailSmtpProvider) Send(subject string, body string, receivers []string, attachments []AttachmentPart) (err error) {
	if len(receivers) == 0 {
		return fmt.Errorf("smtp: no receivers")
	}
	fromHeader := encodeFrom(s.Config.SenderEmail, s.Config.SenderAlias)
	message := buildMessage(fromHeader, subject, body, receivers, attachments)

	addr := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
	auth := smtp.PlainAuth("", s.Config.AuthUser, s.Config.AuthPass, s.Config.Host)
	tlsConfig := &tls.Config{
		ServerName:         s.Config.Host,
		InsecureSkipVerify: s.Config.InsecureSkipVerify,
	}

	if s.Config.Port == 465 {
		return sendSMTPS(addr, s.Config.Host, auth, hasAuth(s.Config), s.Config.SenderEmail, receivers, message, tlsConfig)
	}
	return sendSTARTTLS(addr, s.Config.Host, auth, hasAuth(s.Config), s.Config.SenderEmail, receivers, message, tlsConfig)
}

func hasAuth(c msg.SmtpConfig) bool {
	return c.AuthUser != "" || c.AuthPass != ""
}

func sendSMTPS(addr, host string, auth smtp.Auth, useAuth bool, from string, to []string, message []byte, tlsConfig *tls.Config) error {
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer client.Close()

	return submit(client, auth, useAuth, from, to, message)
}

func sendSTARTTLS(addr, host string, auth smtp.Auth, useAuth bool, from string, to []string, message []byte, tlsConfig *tls.Config) error {
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Close()

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err = client.StartTLS(tlsConfig); err != nil {
			return err
		}
	}

	return submit(client, auth, useAuth, from, to, message)
}

func submit(client *smtp.Client, auth smtp.Auth, useAuth bool, from string, to []string, message []byte) error {
	if useAuth {
		if ok, _ := client.Extension("AUTH"); ok {
			if err := client.Auth(auth); err != nil {
				return err
			}
		}
	}

	if err := client.Mail(from); err != nil {
		return err
	}
	for _, rcpt := range to {
		if err := client.Rcpt(strings.TrimSpace(rcpt)); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	if _, err = w.Write(message); err != nil {
		return err
	}
	if err = w.Close(); err != nil {
		return err
	}
	return client.Quit()
}
