package logger

import (
	"log"
	"os"
	"path"
	"runtime"
	"testing"
)

var Logger *log.Logger

func init() {
	_, ok := os.LookupEnv("DEBUG")
    isTesting := testing.Testing()
	if ok || isTesting {
		//cwd, err := os.Getwd()
        _,  loggerPackageRoot, _, ok := runtime.Caller(0)
		if !ok {
			panic("cannot get logger package root file.")
		}
		logFile, err := os.OpenFile(path.Join(path.Dir(loggerPackageRoot), "../logs/debug.log"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		Logger = log.New(logFile, "", log.Ltime|log.Ldate|log.Lshortfile)
	}
}

