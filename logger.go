package loglevel

import (
	"io"
	"log"
)

// GenericField represents a simple key value field
type GenericField struct {
	Key   string
	Value string
}

// Field returns a filled GenericField
func Field(key, value string) GenericField {
	return GenericField{Key: key, Value: value}
}

// Logger represents an active logging object that generates lines of output to an io.Writer
type Logger struct {
	formatter   Formatter
	innerLogger *log.Logger
	loglevel    int
}

func (l *Logger) validLevel(requestedLoglevel int) bool {
	if l.loglevel == StringToLevel("NOTSET") {
		return false
	}
	return requestedLoglevel >= l.loglevel
}

func (l *Logger) log(logLevel int, msg string, fields ...GenericField) {
	if !l.validLevel(logLevel) {
		return
	}
	formatedMsg := l.formatter.Format(logLevel, msg, fields...)
	l.innerLogger.Println(formatedMsg)
}

// Debug prints log with DEBUG level
func (l *Logger) Debug(msg string, fields ...GenericField) {
	l.log(StringToLevel("DEBUG"), msg, fields...)
}

// Info prints log with INFO level
func (l *Logger) Info(msg string, fields ...GenericField) {
	l.log(StringToLevel("INFO"), msg, fields...)
}

// Warning prints log with WARNING level
func (l *Logger) Warning(msg string, fields ...GenericField) {
	l.log(StringToLevel("WARNING"), msg, fields...)
}

// Error prints log with ERROR level
func (l *Logger) Error(msg string, fields ...GenericField) {
	l.log(StringToLevel("ERROR"), msg, fields...)
}

// Critical prints log with CRITICAL level
func (l *Logger) Critical(msg string, fields ...GenericField) {
	l.log(StringToLevel("CRITICAL"), msg, fields...)
}

// NewLogger returns a pointer of Logger
func NewLogger(out io.Writer, prefix string, flag, logLevel int, formatter Formatter) *Logger {
	innerLogger := log.New(out, prefix, flag)
	return &Logger{loglevel: logLevel, innerLogger: innerLogger, formatter: formatter}
}
