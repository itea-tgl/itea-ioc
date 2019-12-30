package ioc

import (
	config "github.com/itea-tgl/itea-config"
	"reflect"
)

var container *Container

func init ()  {
	container = newContainer()
}

type Ioc struct {
	builder IBuilder
}

func New(opt ...IOption) *Ioc {
	ioc := &Ioc{
		builder: DefaultBuilder(),
	}
	for _, o := range opt {
		o.exec(ioc)
	}
	return ioc
}

// Register class into ioc container for make instance
// Param b maybe class or instance of ioc.Bean
func (i *Ioc) Register(b interface{}) {
	container.set(register(b))
}

// Get instance of t
// Param t maybe class or name(string) of class
func (i *Ioc) Instance(t interface{}) interface{} {
	if ins := container.load(t); ins != nil {
		return ins
	}
	return i.builder.make(reflect.TypeOf(t))
}

type IOption interface {
	exec(*Ioc)
}

type AppendOption func(*Ioc)

func (a AppendOption) exec(i *Ioc) {
	a(i)
}

// With config to enable constant inject
func WithConfig(c *config.Config) IOption {
	return AppendOption(func(i *Ioc) {
		i.builder.load(c)
	})
}