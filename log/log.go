package log

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var DebugFlag bool

func logging(level string, v ...interface{}) {
	_, file, lineno, ok := runtime.Caller(2)
	if !ok {
		panic("logging panic")
	}
	var v2 []interface{}
	fileline := fmt.Sprintf("%s:%d", filepath.Base(file), lineno)
	v2 = append(v2, fileline)
	v2 = append(v2, level)
	v2 = append(v2, v...)
	log.Println(v2)
}

func Debug(v ...interface{}) {
	if DebugFlag {
		logging("[DEBUG]", v...)
	}
}

func Info(v ...interface{}) {
	logging("[INFO]", v...)
}

func Error(v ...interface{}) {
	logging("[ERROR]", v...)
}
