// Package logger provides a basic level logging system
package logger

import "log"

type Level uint32

const (
	DebugLevel = iota
	ErrorLevel
)

type Logger struct {
	Enabled bool
	Level   Level
}

// Creates a new logger. It can be disabled.
// All logs with a level greater or equal to the
// selected level will be printed to the standard
// output.
func New(enabled bool, level Level) *Logger {
	return &Logger{
		Enabled: enabled,
		Level:   level,
	}
}

func (logger *Logger) mustBePrinted(level Level) bool {
	return logger.Enabled && logger.Level <= level
}

func (logger *Logger) Debug(messageArgs ...interface{}) {
	if logger.mustBePrinted(DebugLevel) {
		log.Println("[DEBUG]", messageArgs)
	}
}

func (logger *Logger) Error(messageArgs ...interface{}) {
	if logger.mustBePrinted(ErrorLevel) {
		log.Println("[ERROR]", messageArgs)
	}
}
