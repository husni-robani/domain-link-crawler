package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger interface{
	Info(msg string, fields ... interface{})
	Error(msg string, fields ... interface{})
	Fatal(msg string, fields ... interface{})
}

type LogLevel int

const (
	LevelInfo LogLevel = iota + 1
	LevelError
	LevelFatal
)

type logger struct{
	level LogLevel
	output io.Writer
	mu sync.Mutex
}

func New(level LogLevel, output io.Writer) Logger{
	return &logger{
		level: level,
		output: output,
		mu: sync.Mutex{},
	}
}

func (l *logger) log(level LogLevel, msg string, fields ...interface{}) {
	if level < l.level{
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	levelStr := levelToString(l.level)
	logMsg := fmt.Sprintf("[%s] %s", levelStr, msg)

	if len(fields) > 0 {
		logMsg = fmt.Sprintf("%s %v", logMsg, fields)
	}

	fmt.Fprintln(l.output, logMsg)
}

func (l *logger) Info(msg string, fields ...interface{}) {
	l.log(LevelInfo, msg, fields...)
}

func (l *logger) Error(msg string, fields ...interface{}) {
	l.log(LevelError, msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	defer os.Exit(1)

	l.log(LevelFatal, msg, fields...)
}


func levelToString(level LogLevel) string {
    switch level {
    case LevelInfo:
        return "INFO"
    case LevelError:
        return "ERROR"
    case LevelFatal:
        return "FATAL"
    default:
        return "UNKNOWN"
    }
}

var InfoDefaultLogger = New(LevelInfo, os.Stdout)
var ErrDefaultLogger = New(LevelError, os.Stderr)
var FatalDefaultLogger = New(LevelFatal, os.Stderr)