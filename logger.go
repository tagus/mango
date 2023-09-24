package mango

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
)

type LogLevel int8

var (
	logger *log.Logger = log.New(os.Stdout, "", log.Ltime)
	lvl    LogLevel    = LogLevelDebug
)

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelError:
		return color.RedString("[ERROR]")
	case LogLevelWarning:
		return color.YellowString("[WARN]")
	case LogLevelInfo:
		return color.GreenString("[INFO]")
	case LogLevelDebug:
		return "[DEBUG]"
	default:
		return ""
	}
}

func Init(level LogLevel, prefix string) {
	lvl = level
	lp := ""
	if prefix != "" {
		lp = fmt.Sprintf("[%s] ", prefix)
	}
	logger = log.New(os.Stdout, lp, log.Ltime)
}

func decorate(l LogLevel, format string) string {
	if format == "" {
		return fmt.Sprintf("%v", l)
	}
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	return fmt.Sprintf("%v %s", l, format)
}

func Debug(v ...any) {
	if lvl > LogLevelDebug {
		return
	}
	args := []any{decorate(LogLevelDebug, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Debugf(format string, v ...any) {
	if lvl > LogLevelDebug {
		return
	}
	logger.Printf(decorate(LogLevelDebug, format), v...)
}

func Warning(v ...any) {
	if lvl > LogLevelWarning {
		return
	}
	args := []any{decorate(LogLevelWarning, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Warningf(format string, v ...any) {
	if lvl > LogLevelWarning {
		return
	}
	logger.Printf(decorate(LogLevelWarning, format), v...)
}

func Info(v ...any) {
	if lvl > LogLevelInfo {
		return
	}
	args := []any{decorate(LogLevelInfo, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Infof(format string, v ...any) {
	if lvl > LogLevelInfo {
		return
	}
	logger.Printf(decorate(LogLevelInfo, format), v...)
}

func Error(v ...any) {
	if lvl > LogLevelError {
		return
	}
	args := []any{decorate(LogLevelError, "")}
	args = append(args, v...)
	logger.Println(args...)
}

func Errorf(v ...any) {
	if lvl > LogLevelError {
		return
	}
	logger.Println(v...)
}

func Fatal(v ...any) {
	debug.PrintStack()
	logger.Fatal(v...)
}

func FatalIf(err error, v ...any) {
	if err == nil {
		return
	}
	args := []any{decorate(LogLevelError, ""), " ", err}
	args = append(args, v...)
	Fatal(args...)
}
