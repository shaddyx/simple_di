package tools

import (
	"fmt"
	"reflect"
	"strings"
)

func GetInstanceType(v any) string {
	if !IsRef(v) {
		v = &v
	}
	if IsType(v) {
		return v.(reflect.Type).Elem().Name()
	}
	return reflect.ValueOf(v).Type().Elem().Name()
}

func GetInstancePackage(v any) string {
	if !IsRef(v) {
		v = &v
	}
	var path string
	if IsType(v) {
		path = v.(reflect.Type).Elem().PkgPath()
	} else {
		path = reflect.ValueOf(v).Type().Elem().PkgPath()
	}
	return strings.ReplaceAll(path, "/", ".")
}

func GetInstanceQualifier(v any) string {
	if _, ok := v.(string); ok {
		return v.(string)
	}
	return fmt.Sprintf("%s.%s", GetInstancePackage(v), GetInstanceType(v))
}

func GetInterfaceType[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func GetInterfacePackage[T any]() string {
	path := reflect.TypeOf((*T)(nil)).Elem().PkgPath()
	return strings.ReplaceAll(path, "/", ".")
}

func GetQualifier[T any]() string {
	return fmt.Sprintf("%s.%s", GetInterfacePackage[T](), GetInterfaceType[T]())
}

func GetStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}

func GetFunctionReturnType(fn any, num int) reflect.Type {
	if fn == nil {
		panic("function is nil")
	}
	if !IsFunc(fn) {
		panic("not a function")
	}

	return reflect.Indirect(reflect.ValueOf(fn)).Type().Out(num)
}

func GetReferenceType(ref any) reflect.Type {
	return reflect.Indirect(reflect.ValueOf(ref)).Type()
}

func SetValue(obj any, field string, value any) error {
	ref := reflect.ValueOf(obj)

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// should double check we now have a struct (could still be anything)
	if ref.Kind() != reflect.Struct {
		return fmt.Errorf("value is not a struct")
	}

	prop := ref.FieldByName(field)
	prop.Set(reflect.ValueOf(value))
	return nil
}
