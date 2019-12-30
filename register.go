package ioc

import (
	"reflect"
	"strings"
)

const (
	SINGLETON = "singleton"
	CONSTRUCT = "Construct"
)

// Register bean of class
func register(class interface{}) *Bean {
	if v, ok := class.(*Bean); ok {
		return registerBean(v)
	}

	return registerBean(&Bean{
		Concrete: class,
	})
}

// Register bean
func registerBean(b *Bean) *Bean {
	if b.Concrete == nil {
		panic("concrete of bean should not be nil")
	}

	tc := reflect.TypeOf(b.Concrete)
	b.SetConcreteType(tc)
	if b.Abstract == nil {
		b.Abstract = b.Concrete
	}

	ta := reflect.TypeOf(b.Abstract)
	b.SetAbstractType(ta)

	if strings.EqualFold(b.Name, "") {
		b.Name = ta.Name()
	}

	if strings.EqualFold(b.Scope, "") {
		b.Scope = SINGLETON
	}

	if strings.EqualFold(b.Construct, "") {
		b.Construct = CONSTRUCT
	}
	return b
}