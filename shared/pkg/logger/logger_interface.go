package logger

import (
	"io"
)

type LoggerInterface interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
	SetOutput(out io.Writer)
}
