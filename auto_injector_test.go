package simple_di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ServiceToAutoInjectTest struct {
}
type ServiceToNamedInjects struct {
}
type EmbeddedServiceWithAuto struct {
	S1 *ServiceToAutoInjectTest `inject:"autoinject"`
	S2 *ServiceToAutoInjectTest `inject:"autoinject"`
	S3 *ServiceToNamedInjects   `inject:"serviceNamed"`
}

type ServiceWithAutoInjects struct {
	EmbeddedServiceWithAuto
	S4 *ServiceToAutoInjectTest `inject:"autoinject"`
}

func Test_autoInject(t *testing.T) {
	s1 := &ServiceToAutoInjectTest{}
	s2 := &ServiceWithAutoInjects{}
	sNamed := &ServiceToNamedInjects{}
	container := NewContainer().
		AddByName("serviceNamed", sNamed).
		AddByType(s1).
		AddByName("service2", s2)
	s2Got := container.GetPanic("service2").(*ServiceWithAutoInjects)
	assert.Equal(t, s1, s2Got.S1)
	assert.Equal(t, s1, s2Got.S2)
	assert.Equal(t, sNamed, s2Got.S3)
	assert.Equal(t, s1, s2Got.S4)

}
