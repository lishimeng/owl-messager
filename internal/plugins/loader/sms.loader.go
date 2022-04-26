package loader

type SmsLoader interface {
	Load(id string)
	Unload(id string)
}

func New() (sms SmsLoader) {
	h := smsLoaderImpl{}
	sms = &h
	return
}

type smsLoaderImpl struct {
}

func (s *smsLoaderImpl) Load(id string) {

}

func (s *smsLoaderImpl) Unload(id string) {

}