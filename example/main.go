package main

import (
	"log"
	"os"

	"github.com/allisson/go-loglevel"
)

var (
	simpleLogger *loglevel.Logger
	jsonLogger   *loglevel.Logger
)

func init() {
	level := loglevel.LevelFromEnv()
	simpleLogger = loglevel.NewLogger(os.Stdout, "", log.LstdFlags, level, loglevel.NewSimpleFormatter())
	jsonLogger = loglevel.NewLogger(os.Stdout, "", log.LstdFlags, level, loglevel.NewJSONFormatter())
}

func main() {
	simpleLogger.Info("Info Message", loglevel.Field("key-1", "value-1"), loglevel.Field("key-2", "value-2"))
	// output: 2019/09/23 17:33:02 INFO Info Message key-1=value-1 key-2=value-2
	jsonLogger.Info("Info Message", loglevel.Field("key-1", "value-1"), loglevel.Field("key-2", "value-2"))
	// output: 2019/09/23 17:33:02 INFO Info Message {"key-1":"value-1","key-2":"value-2"}
}
