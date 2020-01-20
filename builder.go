package ioc

import (
	config "github.com/itea-tgl/itea-config"
	"reflect"
)

const (
	WIRED = "wired"
	VALUE = "value"
)

type IBuilder interface {
	load(*config.Config)
	make(reflect.Type) interface{}
}

type Builder struct {
	c *config.Config
}

func (b *Builder) make(c reflect.Type) interface{} {
	bean := container.get(c)
	if bean == nil {
		return nil
	}

	t := bean.GetConcreteType()

	ins := reflect.New(t)

	if m := ins.MethodByName(bean.Construct); m.IsValid() {
		m.Call(nil)
	}

	for i := 0; i < t.NumField(); i++ {
		vf := ins.Elem().Field(i)
		if !vf.CanSet() {
			continue
		}
		if v := b.field(vf.Kind(), t.Field(i)); v.IsValid() {
			vf.Set(v)
		}
	}

	if m := ins.MethodByName(bean.InitFunc); m.IsValid() {
		m.Call(nil)
	}

	if bean.isSingleton() {
		bean.setInstance(ins.Interface())
	}

	return ins.Interface()
}

func (b *Builder) field(kind reflect.Kind, tf reflect.StructField) reflect.Value {
	t := tf.Type
	var v interface{}

	switch kind {
	case reflect.Int, reflect.String, reflect.Bool:
		if tag, ok := check(tf, VALUE); ok && b.c != nil {
			v = b.c.Get(tag)
		}
		break
	case reflect.Struct:
		if _, ok := check(tf, WIRED); ok {
			if i := b.make(t); i != nil {
				return reflect.ValueOf(i).Elem()
			}
		}
		break
	case reflect.Ptr:
		if _, ok := check(tf, WIRED); ok {
			v = b.make(t.Elem())
		}
		break
	default:
		break
	}

	return reflect.ValueOf(v)
}

func (b *Builder) load(c *config.Config) {
	b.c = c
}

func check(field reflect.StructField, tag string) (string, bool) {
	if t := field.Tag.Get(tag); t != "" {
		return t, true
	}
	return "", false
}

func DefaultBuilder() IBuilder {
	return &Builder{}
}
