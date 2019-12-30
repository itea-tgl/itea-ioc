package ioc

import (
	config "github.com/itea-tgl/itea-config"
	"os"
	"reflect"
	"testing"
)

func TestIoc_Register(t *testing.T) {
	type class1 struct {}

	ioc := New()
	ioc.Register(class1{})
	ioc.Register(&Bean{
		Name:         "class",
		Scope:        "singleton",
		Abstract:     nil,
		Concrete:     class1{},
		abstractType: nil,
		concreteType: nil,
	})
}

func TestIoc_Instance_without_params(t *testing.T) {
	type class1 struct {}

	ioc := New()
	ioc.Register(class1{})

	ins1 := ioc.Instance(class1{})
	ins2 := ioc.Instance("class1")

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			&class1{},
			ins1,
		},
		{
			&class1{},
			ins2,
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Ioc_Instance_without_params %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}
}

func TestIoc_Instance_with_params_struct(t *testing.T) {
	type class1 struct {}
	type class2 struct {
		C1 class1	`wired:"true"`
	}

	ioc := New()
	ioc.Register(class1{})
	ioc.Register(class2{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			&class2{C1:class1{}},
			ioc.Instance(class2{}),
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Ioc_Instance_with_params %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}
}

func TestIoc_Instance_with_params_ptr(t *testing.T) {
	type class1 struct {}
	type class2 struct {
		C1 *class1	`wired:"true"`
	}

	ioc := New()
	ioc.Register(class1{})
	ioc.Register(class2{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			&class2{C1:&class1{}},
			ioc.Instance(class2{}),
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Ioc_Instance_with_params %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}
}

func TestIoc_Instance_with_params_const(t *testing.T) {
	type class1 struct {
		V1 string	`value:"config_test.class1.v1"`
		V2 int	`value:"config_test.class1.v2"`
		V3 bool	`value:"config_test.class1.v3"`
	}

	path, _ := os.Getwd()

	conf, _ := config.Init(config.Option{
		File:      path + "/test_config.yml",
		Processor: config.YamlProcessor,
	})

	ioc := New(WithConfig(conf))
	ioc.Register(class1{})

	c1 := &class1{V1:"hello world", V2:123, V3:true}
	c2 := ioc.Instance(class1{}).(*class1)

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			c1.V1,
			c2.V1,
		},
		{
			c1.V2,
			c2.V2,
		},
		{
			c1.V3,
			c2.V3,
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Ioc_Instance_with_params %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}
}