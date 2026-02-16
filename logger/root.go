package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	_"testing"
)

var Logger *log.Logger

func init() {
	//_, ok := os.LookupEnv("DEBUG")
    //isTesting := testing.Testing()
	//if ok || isTesting {
		//cwd, err := os.Getwd()
        _,  loggerPackageRoot, _, ok := runtime.Caller(0)
		if !ok {
			panic("cannot get logger package root file.")
		}
		logFile, err := os.OpenFile(path.Join(path.Dir(loggerPackageRoot), "../logs/debug.log"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		Logger = log.New(logFile, "", 0)
	//}
}

func Log(v ...any) {
	if Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		rootDir, _ := os.Getwd()
		relativeFilePath, _ := filepath.Rel(rootDir, file)
		toLog := []any{fmt.Sprintf("%s:%d: ", relativeFilePath[3:], line)}
		toLog = append(toLog, v...)
		Logger.Print(toLog...)
	}
}
