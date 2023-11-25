package tools

import (
	"reflect"
	"simpledi/sumpledi/types"
)

func IsProvider(v any) bool {
	_, ok := v.(types.Provider)
	return ok
}

func IsFunc(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func IsRef(input any) bool {
	if input == nil {
		return false
	}
	t := reflect.TypeOf(input)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func IsType(input any) bool {
	_, ok := input.(reflect.Type)
	return ok
}

func IsInterface(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Interface
}
