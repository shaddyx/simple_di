package tools

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstanceType(t *testing.T) {
	type Test struct {
	}
	assert.Equal(t, "Test", GetInstanceType(&Test{}))
}

func TestGetInstancePackage(t *testing.T) {
	type Test struct {
	}
	assert.Equal(t, "github.com.shaddyx.simple_di.internal.tools", GetInstancePackage(&Test{}))
}

func TestGetInstanceQualifier(t *testing.T) {
	type Test struct {
	}
	assert.Equal(t, "github.com.shaddyx.simple_di.internal.tools.Test", GetInstanceQualifier(&Test{}))
}

func TestGetInstanceQualifierStruct(t *testing.T) {
	type Test struct {
	}
	res := GetInstanceQualifier(Test{})
	assert.Equal(t, ".", res)
}

func TestIsRef(t *testing.T) {
	type Test struct {
	}
	assert.True(t, IsRef(&Test{}))
	assert.False(t, IsRef(Test{}))
}

type Test2 struct {
}

func TestGetFunctionReturnType(t *testing.T) {
	testFunc := func() (string, int, *Test2) {
		return "Test", 1, &Test2{}
	}

	res := GetFunctionReturnType(testFunc, 0)
	assert.Equal(t, "string", res.Name())

	res = GetFunctionReturnType(testFunc, 1)
	assert.Equal(t, "int", res.Name())

	res = GetFunctionReturnType(testFunc, 2)
	assert.Equal(t, reflect.Ptr, res.Kind())
	assert.Equal(t, "Test2", res.Elem().Name())
}

func TestGetFunctionReturnTypeWithTwoParams(t *testing.T) {
	testFunc := func() (string, error) {
		return "Test", nil
	}

	res := GetFunctionReturnType(testFunc, 0)
	assert.Equal(t, "string", res.Name())
}

func TestGetFunctionReturnTypeWithError(t *testing.T) {
	testFunc := func() error {
		return fmt.Errorf("Test")
	}

	res := GetFunctionReturnType(testFunc, 0)
	assert.Equal(t, "error", res.Name())
}

type Test struct {
}

func (t *Test) String() string {
	return "hello"
}

type Test1 interface {
	String() string
}

func TestGetReferenceType(t *testing.T) {

	ref := &Test{}
	res := GetReferenceType(ref)
	assert.Equal(t, "Test", res.Name())

	var t1 Test1 = &Test{}
	res = GetReferenceType(t1)
	assert.Equal(t, "Test", res.Name())
}
