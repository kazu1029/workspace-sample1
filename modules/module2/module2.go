package moduleA2AAAAAAAAAi

import (
	"fmt"
	"github.com/kazu1029/workspace-sample1/hoge/fuga"
)

type module2Service struct {
}

type Module2Service interface {
	Print(str string)
}

func NewModule2Service() Module2Service {
	return &module2Service{}
}

func (s *module2Service) Print(str string) {
	fmt.Println("module1: ", fuga.Print(str))
}
