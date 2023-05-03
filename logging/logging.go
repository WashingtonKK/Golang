package logging

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	timeFormat string
	debug      bool
}

// The function creates a new logger and returns a pointer to a logger instance
func New(timeFormat string, debug bool) *Logger {
	return &Logger{
		timeFormat: timeFormat,
		debug:      debug,
	}
}

// This function logs and it operates on a logger instanced
func (l *Logger) Log(level string, s string) {
	level = strings.ToLower(level)
	switch level {
	case "info", "warning":
		if l.debug {
			l.write(level, s)
		}
	default:
		l.write(level, s)
	}
}

// This is an unexported method.
func (l *Logger) write(level string, s string) {
	fmt.Printf("[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
}
