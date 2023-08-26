package mango

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type LogLevel int8

var (
	logger *log.Logger
	lvl    LogLevel
)

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelError:
		return "ERROR"
	case LogLevelWarning:
		return "WARN"
	case LogLevelInfo:
		return "INFO"
	case LogLevelDebug:
		return "DEBUG"
	default:
		return ""
	}
}

func Init(level LogLevel, prefix string) {
	lvl = level
	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", prefix), log.Ltime)
}

func check() {
	if logger == nil {
		fmt.Println("logging was never initialized")
		os.Exit(1)
	}
}

func decorate(l LogLevel, format string) string {
	if format == "" {
		return fmt.Sprintf("[%v]", l)
	}
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	return fmt.Sprintf("[%v] %s", l, format)
}

func Debug(v ...any) {
	check()
	if LogLevelDebug > lvl {
		return
	}
	args := []any{decorate(LogLevelDebug, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Debugf(format string, v ...any) {
	check()
	if LogLevelDebug > lvl {
		return
	}
	logger.Printf(decorate(LogLevelDebug, format), v...)
}

func Warning(v ...any) {
	check()
	if LogLevelWarning > lvl {
		return
	}
	args := []any{decorate(LogLevelWarning, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Warningf(format string, v ...any) {
	check()
	if LogLevelWarning > lvl {
		return
	}
	logger.Printf(decorate(LogLevelWarning, format), v...)
}

func Info(v ...any) {
	check()
	if LogLevelInfo > lvl {
		return
	}
	args := []any{decorate(LogLevelInfo, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Infof(format string, v ...any) {
	check()
	if LogLevelInfo > lvl {
		return
	}
	logger.Printf(decorate(LogLevelInfo, format), v...)
}

func Error(v ...any) {
	check()
	if LogLevelError > lvl {
		return
	}
	args := []any{decorate(LogLevelError, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Errorf(v ...any) {
	check()
	if LogLevelError > lvl {
		return
	}
	logger.Println(v...)
}

func Fatal(v ...any) {
	check()
	logger.Fatal(v...)
}

func FatalIf(err error, v ...any) {
	check()
	if err == nil {
		return
	}
	args := []any{decorate(LogLevelError, ""), " ", err}
	args = append(args, v...)
	logger.Fatal(args...)
}
