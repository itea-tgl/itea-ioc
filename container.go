package ioc

import "reflect"

type Container struct {
	bt map[reflect.Type]*Bean
	bn map[string]*Bean
}

func newContainer() *Container {
	return &Container{
		bt: make(map[reflect.Type]*Bean),
		bn: make(map[string]*Bean),
	}
}

// Set bean into container
func (c *Container) set(bean *Bean) {
	c.bt[bean.GetAbstractType()] = bean
	c.bn[bean.Name] = bean
}

// Get bean from container
// Param index maybe class, name(string) of class or reflect type of class
func (c *Container) get(index interface{}) *Bean {
	if v, ok := index.(string); ok {
		if b, ok := c.bn[v]; ok {
			return b
		}
	}

	var v reflect.Type
	if _, ok := index.(reflect.Type); ok {
		v = index.(reflect.Type)
	} else {
		v = reflect.TypeOf(index)
	}

	if b, ok := c.bt[v]; ok {
		return b
	}

	return nil
}

// Get instance of referred class or name(string) of class
func (c *Container) load(i interface{}) interface{} {
	if b := c.get(i); b != nil {
		return b.getInstance()
	}
	return nil
}
