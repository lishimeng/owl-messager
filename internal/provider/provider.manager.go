package provider

type Factory struct {
}

var Default *Factory

func init() {
	Default = &Factory{}
}

func (f *Factory) Create() {

}