package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRef(t *testing.T) {
	type Test struct {
	}
	assert.True(t, IsRef(&Test{}))
	assert.False(t, IsRef(Test{}))
}

func TestIsInterface(t *testing.T) {

	type Test struct {
	}

	type Test1 interface {
		String() string
	}

	assert.False(t, IsInterface(Test{}))
	assert.False(t, IsInterface(&Test{}))
	//assert.True(t, IsInterface(v))
	assert.True(t, IsInterface((*Test1)(nil)))
}
