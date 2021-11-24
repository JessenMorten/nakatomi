package logging

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type consoleLogger struct {
	context  string
	logLevel LogLevel
}

func (c consoleLogger) Debug(format string, a ...interface{}) {
	c.log(Debug, format, a...)
}

func (c consoleLogger) Information(format string, a ...interface{}) {
	c.log(Information, format, a...)
}

func (c consoleLogger) Warning(format string, a ...interface{}) {
	c.log(Warning, format, a...)
}

func (c consoleLogger) Error(format string, a ...interface{}) {
	c.log(Error, format, a...)
}

func (c consoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if c.logLevel <= level {
		levelText := "UNKNOWN"

		if level == Debug {
			levelText = "DEBUG"
		} else if level == Information {
			levelText = "INFORMATION"
		} else if level == Warning {
			levelText = "WARNING"
		} else if level == Error {
			levelText = "ERROR"
		}

		timestamp := time.Now().Format("15:04:05")
		message := fmt.Sprintf(format, a...)
		context := fmt.Sprintf("[%v]", c.context)
		message = fmt.Sprintf("%v %-12s %-15s: %v", timestamp, levelText, context, message)

		if level == Debug {
			color.New(color.FgHiWhite).Println(message)
		} else if level == Information {
			color.New(color.FgHiCyan).Println(message)
		} else if level == Warning {
			color.New(color.FgHiYellow).Println(message)
		} else if level == Error {
			color.New(color.FgHiRed).Println(message)
		} else {
			fmt.Println(message)
		}
	}
}
