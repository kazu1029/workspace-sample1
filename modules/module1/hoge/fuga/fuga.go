package fuga

import (
	"fmt"
	"github.com/kazu1029/workspace-sample1/internal/logging"
)

func Print(str string) string {
	logging.Debug(str)
	return fmt.Sprintf("fuga: %s", str)
}
