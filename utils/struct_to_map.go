package utils

import (
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		dataKey := t.Field(i).Tag.Get("json")
		if dataKey == "" {
			dataKey = SnakeString(t.Field(i).Name)
		}
		dataValue := t.Field(i).Tag.Get("default")
		if dataValue == "" {
			dataValue = v.Field(i).Interface().(string)
		}
		data[dataKey] = dataValue
	}
	return data
}
