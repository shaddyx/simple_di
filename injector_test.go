package simple_di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ServiceToInjectTest struct {
}

type EmbeddedService struct {
	S1 *ServiceToInjectTest `inject:"service1"`
	S2 *ServiceToInjectTest `inject:"service1"`
	S3 *ServiceToInjectTest `inject:"service1"`
}

type ServiceWithInjects struct {
	EmbeddedService
	S4 *ServiceToInjectTest `inject:"service1"`
}

func Test_simpleInject(t *testing.T) {
	s1 := &ServiceToInjectTest{}
	s2 := &ServiceWithInjects{}
	container := NewContainer().
		AddByName("service1", s1).
		AddByName("service2", s2)
	s2Got := container.GetPanic("service2").(*ServiceWithInjects)
	assert.Equal(t, s2Got.S1, s1)
	assert.Equal(t, s2Got.S2, s1)
	assert.Equal(t, s2Got.S3, s1)
	assert.Equal(t, s2Got.S4, s1)

}

func Test_simpleInjectFunctionalInitializer(t *testing.T) {
	s1 := &ServiceToInjectTest{}
	s2 := &ServiceWithInjects{}
	var called bool = false
	container := NewContainer().
		AddFuncByName("service1", func(c *Container) any {
			called = true
			return s1
		}).
		AddByName("service2", s2)
	s2Got := container.GetPanic("service2").(*ServiceWithInjects)
	assert.Equal(t, s2Got.S1, s1)
	assert.Equal(t, s2Got.S2, s1)
	assert.Equal(t, s2Got.S3, s1)
	assert.Equal(t, s2Got.S4, s1)
	assert.True(t, called)
}
