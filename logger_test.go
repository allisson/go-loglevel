package loglevel

import (
	"bytes"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	t.Run("log", func(t *testing.T) {
		tables := []struct {
			loggerLogLevel int
			printLogLevel  int
			expectedResult string
		}{
			{StringToLevel("NOTSET"), StringToLevel("DEBUG"), ""},
			{StringToLevel("NOTSET"), StringToLevel("INFO"), ""},
			{StringToLevel("NOTSET"), StringToLevel("WARNING"), ""},
			{StringToLevel("NOTSET"), StringToLevel("ERROR"), ""},
			{StringToLevel("NOTSET"), StringToLevel("CRITICAL"), ""},
			{StringToLevel("DEBUG"), StringToLevel("DEBUG"), fmt.Sprintf("%s DEBUG log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("DEBUG"), StringToLevel("INFO"), fmt.Sprintf("%s INFO log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("DEBUG"), StringToLevel("WARNING"), fmt.Sprintf("%s WARNING log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("DEBUG"), StringToLevel("ERROR"), fmt.Sprintf("%s ERROR log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("DEBUG"), StringToLevel("CRITICAL"), fmt.Sprintf("%s CRITICAL log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("INFO"), StringToLevel("DEBUG"), ""},
			{StringToLevel("INFO"), StringToLevel("INFO"), fmt.Sprintf("%s INFO log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("INFO"), StringToLevel("WARNING"), fmt.Sprintf("%s WARNING log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("INFO"), StringToLevel("ERROR"), fmt.Sprintf("%s ERROR log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("INFO"), StringToLevel("CRITICAL"), fmt.Sprintf("%s CRITICAL log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("WARNING"), StringToLevel("DEBUG"), ""},
			{StringToLevel("WARNING"), StringToLevel("INFO"), ""},
			{StringToLevel("WARNING"), StringToLevel("WARNING"), fmt.Sprintf("%s WARNING log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("WARNING"), StringToLevel("ERROR"), fmt.Sprintf("%s ERROR log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("WARNING"), StringToLevel("CRITICAL"), fmt.Sprintf("%s CRITICAL log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("ERROR"), StringToLevel("DEBUG"), ""},
			{StringToLevel("ERROR"), StringToLevel("INFO"), ""},
			{StringToLevel("ERROR"), StringToLevel("WARNING"), ""},
			{StringToLevel("ERROR"), StringToLevel("ERROR"), fmt.Sprintf("%s ERROR log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("ERROR"), StringToLevel("CRITICAL"), fmt.Sprintf("%s CRITICAL log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
			{StringToLevel("CRITICAL"), StringToLevel("DEBUG"), ""},
			{StringToLevel("CRITICAL"), StringToLevel("INFO"), ""},
			{StringToLevel("CRITICAL"), StringToLevel("WARNING"), ""},
			{StringToLevel("CRITICAL"), StringToLevel("ERROR"), ""},
			{StringToLevel("CRITICAL"), StringToLevel("CRITICAL"), fmt.Sprintf("%s CRITICAL log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))},
		}
		for _, table := range tables {
			out := bytes.NewBufferString("")
			logger := NewLogger(out, "", log.Ldate, table.loggerLogLevel, NewSimpleFormatter())
			logger.log(table.printLogLevel, "log-msg", Field("key1", "value1"), Field("key2", "value2"))

			if out.String() != table.expectedResult {
				t.Errorf("invalid log result, got=%s, want=%s", out.String(), table.expectedResult)
			}
		}
	})

	t.Run("Debug", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%s DEBUG log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))
		out := bytes.NewBufferString("")
		logger := NewLogger(out, "", log.Ldate, StringToLevel("DEBUG"), NewSimpleFormatter())
		logger.Debug("log-msg", Field("key1", "value1"), Field("key2", "value2"))

		if out.String() != expectedResult {
			t.Errorf("invalid log result, got=%s, want=%s", out.String(), expectedResult)
		}
	})

	t.Run("Info", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%s INFO log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))
		out := bytes.NewBufferString("")
		logger := NewLogger(out, "", log.Ldate, StringToLevel("DEBUG"), NewSimpleFormatter())
		logger.Info("log-msg", Field("key1", "value1"), Field("key2", "value2"))

		if out.String() != expectedResult {
			t.Errorf("invalid log result, got=%s, want=%s", out.String(), expectedResult)
		}
	})

	t.Run("Warning", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%s WARNING log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))
		out := bytes.NewBufferString("")
		logger := NewLogger(out, "", log.Ldate, StringToLevel("DEBUG"), NewSimpleFormatter())
		logger.Warning("log-msg", Field("key1", "value1"), Field("key2", "value2"))

		if out.String() != expectedResult {
			t.Errorf("invalid log result, got=%s, want=%s", out.String(), expectedResult)
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%s ERROR log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))
		out := bytes.NewBufferString("")
		logger := NewLogger(out, "", log.Ldate, StringToLevel("DEBUG"), NewSimpleFormatter())
		logger.Error("log-msg", Field("key1", "value1"), Field("key2", "value2"))

		if out.String() != expectedResult {
			t.Errorf("invalid log result, got=%s, want=%s", out.String(), expectedResult)
		}
	})

	t.Run("Critical", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%s CRITICAL log-msg key1=value1 key2=value2\n", time.Now().Format("2006/01/02"))
		out := bytes.NewBufferString("")
		logger := NewLogger(out, "", log.Ldate, StringToLevel("DEBUG"), NewSimpleFormatter())
		logger.Critical("log-msg", Field("key1", "value1"), Field("key2", "value2"))

		if out.String() != expectedResult {
			t.Errorf("invalid log result, got=%s, want=%s", out.String(), expectedResult)
		}
	})
}
