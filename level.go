package loglevel

import (
	"os"
	"strings"
)

// StringToLevel returns a int based on name of level
func StringToLevel(level string) int {
	switch level {
	case "DEBUG":
		return 10
	case "INFO":
		return 20
	case "WARNING":
		return 30
	case "ERROR":
		return 40
	case "CRITICAL":
		return 50
	}
	return 0
}

// IntToLevel returns a string based on int level
func IntToLevel(level int) string {
	switch level {
	case 50:
		return "CRITICAL"
	case 40:
		return "ERROR"
	case 30:
		return "WARNING"
	case 20:
		return "INFO"
	case 10:
		return "DEBUG"
	}
	return "NOTSET"
}

// LevelFromEnv returns a int based on LOG_LEVEL envvar
func LevelFromEnv() int {
	level := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	validLevels := []string{"NOTSET", "DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}
	for _, l := range validLevels {
		if l == level {
			return StringToLevel(level)
		}
	}
	return StringToLevel("INFO")
}
