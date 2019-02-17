package log

import (
	"fmt"
	"sync"
)

const defaultLogger = "default"

var loggers = make(map[string]Logger)
var mutex = &sync.Mutex{}

// SetDefaultLogger sets the default logger
func SetDefaultLogger(l Logger) {
	SetLogger(defaultLogger, l)
}

// SetLogger sets a named logger
func SetLogger(name string, l Logger) {
	mutex.Lock()
	defer mutex.Unlock()
	loggers[name] = l
}

// GetDefaultLogger returns the default logger
func GetDefaultLogger() (Logger, error) {
	return GetLogger(defaultLogger)
}

// GetLogger returns a named logger
func GetLogger(name string) (Logger, error) {
	if l, ok := loggers[name]; ok {
		return l, nil
	}
	return nil, fmt.Errorf("logger %q not found", name)
}

// Info outputs a message to the log
func Info(msg string, kv ...KV) {
	loggers[defaultLogger].Info(msg, kv...)
}

// Error outputs an error message to the log
func Error(err error, msg string, kv ...KV) {
	loggers[defaultLogger].Error(err, msg, kv...)
}

// V returns a log writter at the specified level
func V(level int) InfoWriter {
	return loggers[defaultLogger].V(level)
}
