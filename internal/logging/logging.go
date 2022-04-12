package logging

import (
	"log"
)

func Debug(str string) {
	log.Println("[DEBUG] ", str)
}

func Info(str string) {
	log.Println("[INFO] ", str)
}
