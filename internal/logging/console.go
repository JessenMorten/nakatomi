package logging

import (
	"fmt"
	"time"
)

type consoleLogger struct {
}

func (c consoleLogger) Debug(format string, a ...interface{}) {
	log("DEBUG", format, a...)
}

func (c consoleLogger) Information(format string, a ...interface{}) {
	log("INFO", format, a...)
}

func (c consoleLogger) Warning(format string, a ...interface{}) {
	log("WARNING", format, a...)
}

func (c consoleLogger) Error(format string, a ...interface{}) {
	log("ERROR", format, a...)
}

func log(level string, format string, a ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, a...)
	message = fmt.Sprintf("%v %v: %v", timestamp, level, message)
	fmt.Println(message)
}
