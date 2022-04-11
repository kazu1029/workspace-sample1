package module2

import (
	"fmt"

	"hoge/fuga"
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
