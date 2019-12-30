package ioc

import "reflect"

type Bean struct {

	//name of bean, use for get instance by name
	Name 			string

	//singleton or prototype
	Scope 			string

	//abstract type of bean
	Abstract 		interface{}

	//concrete type of bean
	Concrete 		interface{}

	//construct function of instance
	//execute when instance be build
	Construct		string

	//function execute after instance be build
	InitFunc		string

	//reflect type of abstract
	abstractType 	reflect.Type

	//reflect type of concrete
	concreteType 	reflect.Type

	//the instance of concrete
	instance 		interface{}
}

func (b *Bean) SetAbstractType(t reflect.Type) {
	b.abstractType = t
}

func (b *Bean) SetConcreteType(t reflect.Type) {
	b.concreteType = t
}

func (b *Bean) GetAbstractType() reflect.Type {
	return b.abstractType
}

func (b *Bean) GetConcreteType() reflect.Type{
	return b.concreteType
}

func (b *Bean) setInstance(i interface{}) {
	b.instance = i
}

func (b *Bean) getInstance() interface{} {
	return b.instance
}

func (b *Bean) isSingleton() bool {
	return b.Scope == SINGLETON
}