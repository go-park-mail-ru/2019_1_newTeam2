package logger

import (
	"io"
	"log"
)

type GoLogger struct {
	logger log.Logger
}

func (stdLogger *GoLogger) Log(v ...interface{}) {
	stdLogger.logger.Print("LOG: ")
	stdLogger.logger.Print(v...)
}

func (stdLogger *GoLogger) Logf(format string, v ...interface{}) {
	stdLogger.logger.Printf(format, v...)
}

func (stdLogger *GoLogger) SetOutput(out io.Writer) {
	stdLogger.logger.SetOutput(out)
}
