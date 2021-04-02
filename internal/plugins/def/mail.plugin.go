package def

type Mail interface {
	OnData() (err error)
}
