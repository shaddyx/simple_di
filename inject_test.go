package simple_di

import (
	"testing"

	"github.com/shaddyx/simple_di/internal/tools"
	"github.com/stretchr/testify/assert"
)

type Service1 struct {
}

type TestStruct struct {
	S1         *Service1 `inject:"service1"`
	S3         *Service1 `inject:"service1"`
	S4         int
	S5         string
	SomeString string `inject:"env:SOME_STRING"`
}

type TestEmbeddedStruct struct {
	TestStruct
	S34 *Service1 `inject:"service1"`
}

func Test_getAllFields(t *testing.T) {
	fields := getAllFields(&TestStruct{})
	assert.Equal(t, 3, len(fields))
	assert.Equal(t, "S1", fields["S1"].Name)
	assert.Equal(t, "S3", fields["S3"].Name)
}

func Test_getAllFieldsEmbedded(t *testing.T) {
	fields := getAllFields(&TestEmbeddedStruct{})
	assert.Equal(t, 4, len(fields))
	assert.Equal(t, "S34", fields["S34"].Name)
	assert.Equal(t, "S3", fields["S3"].Name)
	assert.Equal(t, "S1", fields["S1"].Name)
	assert.Equal(t, "SomeString", fields["SomeString"].Name)
}

func TestIsFunc(t *testing.T) {
	testFunc := func() {

	}

	assert.True(t, tools.IsFunc(testFunc))
	assert.False(t, tools.IsFunc(1))
	assert.False(t, tools.IsFunc(Service1{}))
	assert.False(t, tools.IsFunc(&Service1{}))
}
