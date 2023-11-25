package simple_di

import (
	"fmt"
	"testing"

	"github.com/shaddyx/simple_di/internal/tools"
	"github.com/stretchr/testify/assert"
)

type InterfaceToInject interface {
	SomeMethod()
}

type StructToBeInjected struct {
}

func (s *StructToBeInjected) SomeMethod() {

}

type StructToInjectInterfaceInTest struct {
	Injected InterfaceToInject `inject:"autoinject"`
}

func Test_interfaceInjection(t *testing.T) {
	s1 := &StructToBeInjected{}
	fmt.Println(tools.GetQualifier[InterfaceToInject]())
	container := NewContainer().
		AddByName(tools.GetQualifier[InterfaceToInject](), s1).
		AddByName("s1", &StructToInjectInterfaceInTest{})
	s := container.GetPanic("s1").(*StructToInjectInterfaceInTest)
	assert.Equal(t, s1, s.Injected)

}
