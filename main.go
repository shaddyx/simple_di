package main

import (
	"reflect"
	"simpledi/internal/tools"
	"simpledi/some/package/somewhere"
)

type ITest interface {
}

type Stest struct {
}

func main() {
	println("type:" + tools.GetInstanceQualifier((*somewhere.ITest1)(nil)))
	println("type:" + tools.GetInstanceQualifier(&Stest{}))
	println("type:" + tools.GetQualifier[somewhere.ITest1]())
	println("type:" + reflect.ValueOf((*ITest)(nil)).Type().Elem().Name())
}
