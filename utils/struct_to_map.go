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
		dataValue := v.Field(i).Interface()
		// TODO fix value type
		if dataValue == "" {
			dataValue = t.Field(i).Tag.Get("default")
		}
		data[dataKey] = dataValue
	}
	return data
}
