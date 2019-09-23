package loglevel

import (
	"os"
	"testing"
)

func TestStringToLevel(t *testing.T) {
	tables := []struct {
		level         string
		expectedLevel int
	}{
		{"NOTSET", 0},
		{"DEBUG", 10},
		{"INFO", 20},
		{"WARNING", 30},
		{"ERROR", 40},
		{"CRITICAL", 50},
	}
	for _, table := range tables {
		result := StringToLevel(table.level)

		if table.expectedLevel != result {
			t.Errorf("invalid level result, got=%d, want=%d", result, table.expectedLevel)
		}
	}
}

func TestIntToLevel(t *testing.T) {
	tables := []struct {
		level         int
		expectedLevel string
	}{
		{0, "NOTSET"},
		{10, "DEBUG"},
		{20, "INFO"},
		{30, "WARNING"},
		{40, "ERROR"},
		{50, "CRITICAL"},
	}
	for _, table := range tables {
		result := IntToLevel(table.level)

		if table.expectedLevel != result {
			t.Errorf("invalid level result, got=%s, want=%s", result, table.expectedLevel)
		}
	}
}

func TestLevelFromEnv(t *testing.T) {
	tables := []struct {
		level         string
		expectedLevel int
	}{
		{"", 20},
		{"INVALID-LEVEL", 20},
		{"NOTSET", 0},
		{"DEBUG", 10},
		{"INFO", 20},
		{"WARNING", 30},
		{"ERROR", 40},
		{"CRITICAL", 50},
	}
	for _, table := range tables {
		os.Setenv("LOG_LEVEL", table.level)
		result := LevelFromEnv()

		if table.expectedLevel != result {
			t.Errorf("invalid level result, got=%d, want=%d", result, table.expectedLevel)
		}
	}
}
