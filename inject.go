package simple_di

import (
	"log"
	"maps"
	"reflect"

	"github.com/shaddyx/simple_di/internal/tools"
)

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

func inject(object any, container *Container) error {
	fields := getAllFields(object)
	log.Println("Got fields: ", fields)
	for name, field := range fields {
		tag := field.Tag.Get("inject")
		if tag == "autoinject" {
			tag = tools.GetInstanceQualifier(field.Type)
		}
		log.Printf("injecting %s -> %s", tag, name)
		f, err := container.GetByName(tag)
		if err != nil {
			return err
		}
		err = tools.SetValue(object, field.Name, f)
		if err != nil {
			return err
		}
	}
	return nil
}
