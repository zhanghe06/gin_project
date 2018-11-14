package utils

import (
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)


func TestRegisterType(t *testing.T) {
	var typeRegistry = TypeRegistry{}
	type Person struct {
		Name string
		Sex  int
	}

	// 注册类型
	typeRegistry.RegisterType((*Person)(nil))

	// 根据名称创建结构
	structName := "Person"
	testStruct, ok := typeRegistry.NewStruct(structName)
	if !ok {
		return
	}
	assert.Equal(t, testStruct, Person{})
	assert.Equal(t, reflect.TypeOf(testStruct).Name(), structName)
}
