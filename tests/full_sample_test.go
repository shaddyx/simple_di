package tests

import (
	"fmt"
	"simpledi/internal/tools"
	simpledi "simpledi/sumpledi"
	"simpledi/sumpledi/types"
	"testing"
)

type Service1 interface {
	String() string
}

type Service2 interface {
	String() string
}

type ServiceImpl1 struct {
	s2 Service2
}

func (s *ServiceImpl1) String() string {
	return "1"
}

// s2 getter
func (s *ServiceImpl1) GetS2() Service2 {
	return s.s2
}

type ServiceImpl2 struct {
	s1 Service1
}

func (s *ServiceImpl2) String() string {
	return "2"
}

func TestFullSample(t *testing.T) {
	container := simpledi.NewContainer()
	container.Register(types.Provider{
		Initializer: &ServiceImpl1{},
	}, types.Provider{
		Initializer: func(s1 Service1) Service2 {
			return &ServiceImpl2{
				s1: s1,
			}
		},
	},
	)
	i, err := container.GetInstance(tools.GetQualifier[ServiceImpl2]())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(i)
}
