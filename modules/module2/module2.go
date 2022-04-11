package module2

import (
	"fmt"
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
	fmt.Println("module2 updated: ", print(str))
}
