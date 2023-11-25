package simple_di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainer_IterateInstances(t *testing.T) {
	s1 := &ServiceToInjectTest{}
	s2 := &ServiceWithInjects{}
	container := NewContainer().
		AddByName("service1", s1).
		AddByName("service2", s2)
	called := make([]string, 0)
	container.IterateInstances(func(name string, obj any) error {
		if name == "service1" {
			assert.Equal(t, s1, obj)
		}
		if name == "service2" {
			assert.Equal(t, s2, obj)
		}
		called = append(called, name)
		return nil
	})
	assert.Equal(t, 2, len(called))
	assert.Equal(t, []string{"service1", "service2"}, called)
}
