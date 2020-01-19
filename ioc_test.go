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

type class struct {
	A string
	B string
}

func (c *class) Construct() {
	c.A = "123"
	c.B = "aaa"
}

func (c *class) Execute() {
	c.B = "bbb"
}

func TestIoc_Instance_with_construct(t *testing.T) {
	ioc := New()
	ioc.Register(&Bean{
		Concrete:     class{},
		InitFunc:     "Execute",
	})

	ins1 := ioc.Instance(class{})
	ins2 := ioc.Instance("class")

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			&class{A:"123", B:"bbb"},
			ins1,
		},
		{
			&class{A:"123", B:"bbb"},
			ins2,
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("TestIoc_Instance_with_construct %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}
}

func TestIoc_Instance_with_params_struct(t *testing.T) {
	type class1 struct {}
	type class2 struct {
		C1 class1	`wired:"true"`
		C2 class1
	}

	ioc := New()
	ioc.Register(class1{})
	ioc.Register(class2{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			&class2{},
			ioc.Instance(class2{}),
		},
		{
			&class2{C1:class1{}},
			ioc.Instance("class2"),
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
	type class2 struct {}
	type class3 struct {
		C1 *class1	`wired:"true"`
		CC1 *class1
		C2 *class2	`wired:"true"`
		cc1 *class1	`wired:"true"`
	}

	ioc := New()
	ioc.Register(class1{})
	ioc.Register(class3{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			&class3{C1:&class1{}},
			ioc.Instance(class3{}),
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
		V1 string	`value:"test_config.class1.v1"`
		V2 int	`value:"test_config.class1.v2"`
		V3 bool	`value:"test_config.class1.v3"`
		V4 int64	`value:"test_config.class1.v2"`
	}

	path, _ := os.Getwd()

	conf, _ := config.Init(config.Option{
		File:      path + "/test_config.yml",
		Processor: config.YamlProcessor,
	})

	ioc := New(WithConfig(conf))
	ioc.Register(class1{})

	c1 := &class1{V1:"hello world", V2:123, V3:true, V4:0}
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
		{
			c1.V4,
			c2.V4,
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Ioc_Instance_with_params %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}
}