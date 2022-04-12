package logging

import (
	"log"
)

func Debug(str string) {
	log.Debug("[DEBUG] ", str)
}

func Info(str string) {
	log.Info("[INFO] ", str)
}
