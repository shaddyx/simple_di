package simple_di

import (
	"fmt"
	"log"
	"maps"
	"reflect"
)

func IsFunc(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func getAllFields(object any) map[string]reflect.StructField {
	res := make(map[string]reflect.StructField)
	var t reflect.Type
	if objectField, ok := object.(reflect.StructField); ok {
		t = objectField.Type
	} else {
		t = reflect.Indirect(reflect.ValueOf(object)).Type()
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			maps.Copy(res, getAllFields(f))
		}
		toInject := f.Tag.Get("inject")
		if toInject != "" {
			res[f.Name] = f
		}

	}
	return res
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

func inject(object any, container *Container) error {
	fields := getAllFields(object)
	log.Println("Got fields: ", fields)
	for name, field := range fields {
		tag := field.Tag.Get("inject")

		log.Printf("injecting %s -> %s", tag, name)
		f, err := container.GetByName(tag)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = SetValue(object, field.Name, f)
		if err != nil {
			return err
		}
	}
	return nil
}
