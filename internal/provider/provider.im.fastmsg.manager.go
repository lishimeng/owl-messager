package provider

import (
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/container"
)

type ImFactory struct {
}

func init() {
	container.Add(&ImFactory{})
}

func GetImFactory() (f *ImFactory) {
	_ = container.Get(f)
	return
}

func (f *ImFactory) Create(vendor msg.MessageProvider, config string) (p messager.ImProvider, err error) {

	log.Info("create im provider")
	b, ok := imProviderBuilders[vendor]
	if !ok {
		err = errors.New("unknown im vendor")
		return
	}
	p, err = b(config)

	return
}

var imProviderBuilders = map[msg.MessageProvider]func(config string) (messager.ImProvider, error){}

func RegisterImProvider(vendor msg.MessageProvider, h func(config string) (messager.ImProvider, error)) {
	if h == nil {
		return
	}
	imProviderBuilders[vendor] = h
}

func GetImProviders() (list []msg.MessageProvider) {
	for p := range imProviderBuilders {
		list = append(list, p)
	}
	return
}
