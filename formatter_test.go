package loglevel

import (
	"testing"
)

func TestSimpleFormatter(t *testing.T) {
	logLevel := StringToLevel("DEBUG")
	formatter := NewSimpleFormatter()
	expectedString := "DEBUG simple-formatter key1=value1 key2=value2"
	resultString := formatter.Format(logLevel, "simple-formatter", Field("key1", "value1"), Field("key2", "value2"))

	if expectedString != resultString {
		t.Errorf("invalid formatter result, got=%s, want=%s", resultString, expectedString)
	}
}

func TestJSONFormatter(t *testing.T) {
	logLevel := StringToLevel("DEBUG")
	formatter := NewJSONFormatter()
	expectedString := `DEBUG json-formatter {"key1":"value1","key2":"value2"}`
	resultString := formatter.Format(logLevel, "json-formatter", Field("key1", "value1"), Field("key2", "value2"))

	if expectedString != resultString {
		t.Errorf("invalid formatter result, got=%s, want=%s", resultString, expectedString)
	}
}
