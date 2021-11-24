package logging

type LogLevel int64

const (
	Debug LogLevel = iota
	Information
	Warning
	Error
)

type Logger interface {
	Debug(format string, a ...interface{})
	Information(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
}

func NewConsoleLogger(logLevel LogLevel, context string) Logger {
	return consoleLogger{
		logLevel: logLevel,
		context:  context,
	}
}
