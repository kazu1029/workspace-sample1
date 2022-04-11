package module1

import (
	"fmt"
	"github.com/kazu1029/workspace-sample1/modules/module1/hoge/fuga"
)

type module1Service struct {
}

type Module1Service interface {
	Print(str string)
}

func NewModule1Service() Module1Service {
	return &module1Service{}
}

func (s *module1Service) Print(str string) {
	fmt.Println("module1: ", fuga.Print(str))
}
