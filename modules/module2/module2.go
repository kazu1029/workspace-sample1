package module2

import (
	"fmt"
	"github.com/kazu1029/workspace-sample1/internal/logging"
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
	logging.Info("module2")
	fmt.Println("module2 updated: ", print(str))
}
