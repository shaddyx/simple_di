package simple_di

import (
	"fmt"

	"github.com/shaddyx/simple_di/internal/tools"
)

type ObjectInstance struct {
	Obj      any
	Injected bool
}
type Container struct {
	initializerMap map[string]func(*Container) any
	instanceMap    map[string]*ObjectInstance
}

func NewContainer() *Container {
	return &Container{
		initializerMap: make(map[string]func(*Container) any),
		instanceMap:    make(map[string]*ObjectInstance),
	}
}

func (c *Container) GetInstances() map[string]any {
	res := make(map[string]any)
	for k, v := range c.instanceMap {
		res[k] = v.Obj
	}
	return res
}

func (c *Container) IterateInstances(f func(string, any) error) error {
	for k := range c.instanceMap {
		instance, err := c.GetByName(k)
		if err != nil {
			return err
		}
		if err = f(k, instance); err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) AddFuncByName(name string, initializer func(*Container) any) *Container {
	c.initializerMap[name] = initializer
	return c
}
func (c *Container) AddByType(service ...any) *Container {
	for _, s := range service {
		c.AddByName(tools.GetInstanceQualifier(s), s)
	}
	return c
}

func (c *Container) AddByName(name string, service any) *Container {
	if tools.IsFunc(service) {
		panic("Please use AddFuncByName for func initializers instead")
	}

	if i, ok := service.(*ObjectInstance); ok {
		c.instanceMap[name] = i
		return c
	}

	if i, ok := service.(ObjectInstance); ok {
		c.instanceMap[name] = &i
		return c
	}

	c.instanceMap[name] = &ObjectInstance{
		service,
		false,
	}
	return c
}

func (c *Container) GetByName(name string) (any, error) {
	s, ok := c.instanceMap[name]

	if !ok {
		initializer, ok := c.initializerMap[name]
		if !ok {
			return nil, fmt.Errorf("service %s not found", name)
		}

		c.AddByName(name, initializer(c))
		s = c.instanceMap[name]
	}

	if s.Injected {
		return s.Obj, nil
	}

	err := inject(s.Obj, c)

	if err != nil {
		return nil, err
	}

	return s.Obj, nil
}

func (c *Container) GetPanic(name string) any {
	s, err := c.GetByName(name)
	if err != nil {
		panic(err)
	}
	return s
}
