package validators

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"regexp"
)

const titleRegexString = "^[\u4E00-\u9FA5A-Za-z0-9_]{2,20}$" // 2~20字符，支持英文大小写、数字、中文以及下划线

var titleRegex = regexp.MustCompile(titleRegexString)

func ValidateTitle(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(string); ok {
		return titleRegex.MatchString(date)
	}
	return false
}
