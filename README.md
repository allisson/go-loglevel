# go-loglevel
[![Build Status](https://travis-ci.org/allisson/go-loglevel.svg)](https://travis-ci.org/allisson/go-loglevel) [![Go Report Card](https://goreportcard.com/badge/github.com/allisson/go-loglevel)](https://goreportcard.com/report/github.com/allisson/go-loglevel)

Simple implementation of logs with level filter.

## About

This is a super simple implementation of logs with level filter using only the stdlib.

If you need something more advanced, I recommend these other libs:

- [zap](https://github.com/uber-go/zap)
- [logrus](https://github.com/sirupsen/logrus)

## How to use

```go
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
    // Get level from LOG_LEVEL envvar
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
```

```bash
go run example/main.go
2019/09/23 17:50:01 INFO Info Message key-1=value-1 key-2=value-2
2019/09/23 17:50:01 INFO Info Message {"key-1":"value-1","key-2":"value-2"}
```

```bash
LOG_LEVEL=CRITICAL go run example/main.go # don't show anything

```
