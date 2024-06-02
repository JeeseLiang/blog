package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

// 日志分级
type Level int8

type Fields map[string]interface{}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case 1:
		return "debug"
	case 2:
		return "info"
	case 3:
		return "warn"
	case 4:
		return "error"
	case 5:
		return "fatal"
	case 6:
		return "panic"
	}
	return ""
}

func NewLogger(prefix string, flag int, w io.Writer) *Logger {
	return &Logger{
		newLogger: log.New(w, prefix, flag),
	}
}

func (l *Logger) copy() *Logger {
	newl := *l
	return &newl
}

// 添加Fields
func (l *Logger) WithFields(f Fields) *Logger {
	newl := l.copy()

	if newl.fields == nil {
		newl.fields = make(Fields)
	}

	for k, v := range f {
		newl.fields[k] = v
	}
	return newl
}

func (l *Logger) WithCaller(s int) *Logger {
	newl := l.copy()
	// 获取函数Caller的调用信息
	pc, file, line, ok := runtime.Caller(s)
	if ok {
		f := runtime.FuncForPC(pc)
		newl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return newl
}

// 堆栈调用信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth, minCallerDepth := 25, 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	//输出信息
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Func.Name())
		callers = append(callers, s)
		if !more {
			break
		}
	}
	newl := l.copy()
	newl.callers = callers
	return newl
}

// 日志格式化
func (l *Logger) JsonFormat(level Level, msg string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = msg
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) Output(level Level, msg string) {
	body, _ := json.Marshal(l.JsonFormat(level, msg))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	// 下面的错误等级较高
	case LevelError:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(f string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(f, v...))
}

func (l *Logger) Infof(f string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(f, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}
