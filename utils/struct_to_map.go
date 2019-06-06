package utils

import (
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		// 获取json key
		dataKey := t.Field(i).Tag.Get("json")
		if dataKey == "" {
			dataKey = SnakeString(t.Field(i).Name)
		}

		dataValue := v.Field(i).Interface()

		// 获取默认值
		defaultValue := t.Field(i).Tag.Get("default")
		if defaultValue == "" {
			dataValue = v.Field(i).Interface().(string)
		}

		data[dataKey] = dataValue
	}
	return data
}
