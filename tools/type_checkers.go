package tools

import (
	"reflect"
)

func IsFunctionalProvider(v any) bool {
	return IsFunc(v)
}

func IsFunc(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func IsRef(input any) bool {
	t := reflect.TypeOf(input)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func IsType(input any) bool {
	_, ok := input.(reflect.Type)
	return ok
}
