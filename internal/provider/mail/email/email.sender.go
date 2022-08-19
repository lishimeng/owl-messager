package email

type Sender interface {
	Send(subject string, body string, to ...string) error
}
