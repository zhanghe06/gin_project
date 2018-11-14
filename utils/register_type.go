package utils

import "reflect"

type TypeRegistry map[string]reflect.Type

func (typeRegistry TypeRegistry) RegisterType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func (typeRegistry TypeRegistry) NewStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}
