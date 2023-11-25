package simpledi

import (
	"fmt"
	"log"
	"reflect"
	"simpledi/internal/tools"
	"simpledi/sumpledi/types"
)

type Container struct {
	providers map[string]types.Provider
	instances map[string]any
}

func NewContainer() *Container {
	return &Container{
		providers: make(map[string]types.Provider),
		instances: make(map[string]any),
	}
}

// init container
func (c *Container) Init() {
	if c.providers == nil {
		c.providers = make(map[string]types.Provider)
	}
	if c.instances == nil {
		c.instances = make(map[string]any)
	}
}

func (c *Container) Register(providers ...types.Provider) error {
	c.Init()
	for _, p := range providers {
		err := c.registerProvider(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) registerProvider(r types.Provider) error {
	c.Init()
	if r.Initializer == nil {
		return fmt.Errorf("provider [%s] should have the initializer: %v+", r.Name, r)
	}

	name, err := tryToGetTheName(r)
	if err != nil {
		return err
	}
	r.Name = name
	if _, ok := c.providers[r.Name]; ok {
		return fmt.Errorf("provider already exists: %s", r.Name)
	}
	c.providers[r.Name] = r
	if !tools.IsFunc(r.Initializer) {
		c.instances[r.Name] = r.Initializer
	}
	return nil
}

func (c *Container) Get(qualifier any) (types.Provider, error) {
	c.Init()
	name := tools.GetInstanceQualifier(qualifier)
	prov, ok := c.providers[name]
	if !ok {
		return types.Provider{}, fmt.Errorf("provider not found: %s", name)
	}
	return prov, nil
}

func (c *Container) GetInstance(qualifier any) (any, error) {
	c.Init()
	name := tools.GetInstanceQualifier(qualifier)
	prov, ok := c.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider not found: %s", name)
	}
	instance, ok := c.instances[name]
	if !ok {
		if tools.IsFunc(prov.Initializer) {
			log.Println("calling function initializer for: ", name)
			//return prov.Initializer, nil
			fnValue := reflect.ValueOf(prov.Initializer)
			args := make([]reflect.Value, 0)
			fnResults := fnValue.Call(args)
			instance = fnResults[0].Interface()
			c.instances[name] = instance
		} else {
			return nil, fmt.Errorf("instance not found: %s", name)
		}
	}
	return instance, nil
}

func tryToGetTheName(r types.Provider) (string, error) {
	if r.Name != "" {
		return r.Name, nil
	}

	if tools.IsFunc(r.Initializer) {
		t := tools.GetFunctionReturnType(r.Initializer, 0)
		if t.Kind() != reflect.Ptr && t.Kind() != reflect.Interface {
			return "", fmt.Errorf("initializer should be a pointer or an interface: %v", r.Initializer)
		}
		name := tools.GetInstanceQualifier(t)
		log.Println("provider name: ", name)
		return name, nil
	}

	if r.Initializer != nil {
		return tools.GetInstanceQualifier(r.Initializer), nil
	}

	return "", fmt.Errorf("initializer is nil: %v", r)
}
