package logging

type Logger interface {
	Debug(format string, a ...interface{})
	Information(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
}

func NewConsoleLogger() Logger {
	return consoleLogger{}
}
