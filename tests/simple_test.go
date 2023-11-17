package tests

import (
	"simpledi/internal/tools"
	simpledi "simpledi/sumpledi"
	"simpledi/sumpledi/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	V int
}

func TestGetInstance(t *testing.T) {
	container := simpledi.NewContainer()
	container.Register(types.Provider{
		Initializer: &Test{},
	})
	qualifier := tools.GetInstanceQualifier(&Test{})
	res, err := container.GetInstance(qualifier)
	assert.NoError(t, err)
	assert.Equal(t, &Test{}, res)
}

func TestGetProvider(t *testing.T) {
	container := simpledi.NewContainer()
	container.Register(types.Provider{
		Initializer: &Test{},
	})
	qualifier := tools.GetInstanceQualifier(&Test{})
	res, err := container.Get(qualifier)
	assert.NoError(t, err)
	assert.Equal(t, &Test{}, res.Initializer)
}

func TestGetInstanceByName(t *testing.T) {
	container := simpledi.NewContainer()
	container.Register(types.Provider{
		Initializer: &Test{
			V: 1,
		},
		Name: "SomeName",
	}, types.Provider{
		Initializer: &Test{
			V: 2,
		},
		Name: "SomeOtherName",
	})
	qualifier := tools.GetInstanceQualifier(&Test{})

	_, err := container.GetInstance(qualifier)
	assert.Error(t, err)

	res, err := container.GetInstance("SomeName")
	assert.NoError(t, err)
	assert.Equal(t, &Test{
		V: 1,
	}, res)

	res, err = container.GetInstance("SomeOtherName")
	assert.NoError(t, err)
	assert.Equal(t, &Test{
		V: 2,
	}, res)
}

func TestGetFunctionalProvider(t *testing.T) {
	container := simpledi.NewContainer()
	container.Register(types.Provider{
		Initializer: func() *Test {
			return &Test{
				V: 1,
			}
		},
		Name: "SomeName",
	}, types.Provider{
		Initializer: func() *Test {
			return &Test{
				V: 2,
			}
		},
	})
	qualifier := tools.GetInstanceQualifier(&Test{})

	instance, err := container.GetInstance(qualifier)
	assert.NoError(t, err)
	assert.Equal(t, &Test{
		V: 2,
	}, instance)

}
