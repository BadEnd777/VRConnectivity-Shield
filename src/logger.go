package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	SuccessLevel
)

type Logger struct {
	enableColors bool
}

// NewLogger creates a new logger instance.
func NewLogger(enableColors bool) *Logger {
	return &Logger{enableColors: enableColors}
}

// GetLogger creates a new logger instance with colors enabled.
func GetLogger() *Logger {
	return NewLogger(true)
}

func (l *Logger) log(level LogLevel, colorFunc func(a ...interface{}) string, message ...interface{}) {
	if l.enableColors {
		message = []interface{}{colorFunc(fmt.Sprint(message...))}
	}
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), strings.TrimRight(fmt.Sprint(message...), "\n"))
}

func (l *Logger) Debug(message ...interface{}) {
	l.log(DebugLevel, color.New(color.FgWhite).SprintFunc(), message...)
}

func (l *Logger) Info(message ...interface{}) {
	l.log(InfoLevel, color.New(color.FgCyan).SprintFunc(), message...)
}

func (l *Logger) Warn(message ...interface{}) {
	l.log(WarnLevel, color.New(color.FgYellow).SprintFunc(), message...)
}

func (l *Logger) Error(message ...interface{}) {
	l.log(ErrorLevel, color.New(color.FgRed).SprintFunc(), message...)
}

func (l *Logger) Fatal(message ...interface{}) {
	l.log(FatalLevel, color.New(color.FgRed).SprintFunc(), message...)
	os.Exit(1) // Note: Fatal will exit the program
}

func (l *Logger) Success(message ...interface{}) {
	l.log(SuccessLevel, color.New(color.FgGreen).SprintFunc(), message...)
}

func (l *Logger) Prompt(message string) string {
	fmt.Printf("[%s] %s", time.Now().Format("15:04:05"), message)
	var input string
	fmt.Scanln(&input)
	return input
}

func (l *Logger) SetEnableColors(enableColors bool) {
	l.enableColors = enableColors
}
