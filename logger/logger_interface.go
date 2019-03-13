package logger

import (
	"io"
)

type LoggerInterface interface {
	Log(v ...interface{})
	SetOutput(out io.Writer)
}