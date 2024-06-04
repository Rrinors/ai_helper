package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// LogLevel defines the level of logging.
type LogLevel int

const (
	DebugLevel LogLevel = iota // DebugLevel logs are typically voluminous, and are usually disabled in production.
	InfoLevel                  // InfoLevel is the default logging priority.
	ErrorLevel                 // ErrorLevel is high priority.
	FatalLevel                 // FatalLevel is used to log fatal errors, which will cause the program to exit.
)

var (
	currentLevel LogLevel       // currentLevel holds the current logging level
	lock         sync.Mutex     // lock is used to ensure thread-safe access to the logger
)

// SetLevel sets the logging level.
func SetLevel(level LogLevel) {
	lock.Lock()
	defer lock.Unlock()
	currentLevel = level
}

// Debug logs a message at level Debug.
func Debug(format string, v ...interface{}) {
	log(DebugLevel, "DEBUG", format, v...)
}

// Info logs a message at level Info.
func Info(format string, v ...interface{}) {
	log(InfoLevel, "INFO", format, v...)
}

// Error logs a message at level Error.
func Error(format string, v ...interface{}) {
	log(ErrorLevel, "ERROR", format, v...)
}

// Fatal logs a message at level Fatal and exits the program.
func Fatal(format string, v ...interface{}) {
	log(FatalLevel, "FATAL", format, v...)
	os.Exit(1)
}

// log outputs a formatted log entry if the level is appropriate.
func log(level LogLevel, label, format string, v ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	if currentLevel <= level {
		pc, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Fprintf(os.Stdout, "%s [%s] (%s:%d %s) %s\n", time.Now().Format("2006-01-02 15:04:05"), label, file, line, funcName, fmt.Sprintf(format, v...))
	}
}
