package msg

type Sender interface {
	Send()
	// 进入空闲
	OnIdle()
}
