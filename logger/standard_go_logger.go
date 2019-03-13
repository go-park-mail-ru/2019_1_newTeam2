package logger

import (
	"log"
	"io"
)

type GoLogger struct {
	logger log.Logger
}

func (stdLogger *GoLogger) Log(v ...interface{}) {
	stdLogger.logger.Print(v)
}

func (stdLogger *GoLogger) SetOutput(out io.Writer){
	stdLogger.logger.SetOutput(out)
}