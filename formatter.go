package loglevel

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Formatter is an interface that handles message format
type Formatter interface {
	Format(logLevel int, msg string, fields ...GenericField) string
}

// SimpleFormatter implements Formatter interface
type SimpleFormatter struct{}

// Format is a implementation of Formatter.Format
func (s *SimpleFormatter) Format(logLevel int, msg string, fields ...GenericField) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %s", IntToLevel(logLevel), msg))
	for _, field := range fields {
		sb.WriteString(fmt.Sprintf(" %s=%s", field.Key, field.Value))
	}
	return sb.String()
}

// NewSimpleFormatter returns a pointer of SimpleFormatter
func NewSimpleFormatter() *SimpleFormatter {
	return &SimpleFormatter{}
}

// JSONFormatter implements Formatter interface
type JSONFormatter struct{}

// Format is a implementation of Formatter.Format
func (j *JSONFormatter) Format(logLevel int, msg string, fields ...GenericField) string {
	var sb strings.Builder
	m := make(map[string]string)
	sb.WriteString(fmt.Sprintf("%s %s", IntToLevel(logLevel), msg))
	for _, field := range fields {
		m[field.Key] = field.Value
	}
	b, _ := json.Marshal(m)
	sb.WriteString(fmt.Sprintf(" %s", string(b)))
	return sb.String()
}

// NewJSONFormatter returns a pointer of JSONFormatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}
