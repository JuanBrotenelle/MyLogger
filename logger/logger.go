package logger

import (
	"fmt"
	"log"
	"time"
)

type Logger struct {
	level string
}

const (
	Reset  = "\033[0m"
	Blue   = "\033[34m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
)

func NewLogger(level string) *Logger {
	return &Logger{
		level: level,
	}
}

func (l *Logger) shouldLog(level string) bool {
	levels := map[string]int{
		"INFO":  1,
		"ERROR": 2,
		"WARN":  3,
		"DEBUG": 4,
	}

	return levels[level] <= levels[l.level]
}

func (l *Logger) logMessage(level, color, message string) {
	if l.shouldLog(level) {
		timestamp := time.Now().Format("15:04:05")
		fmt.Println("")
		fmt.Printf("%s[%s]%s %s - %s\n", color, level, Reset, timestamp, message)
		fmt.Println("")
	}
}

/*
If there is a debug, the message will be printed in GREEN
*/
func (l *Logger) Debug(message string) {
	l.logMessage("DEBUG", Green, message)
}

/*
If there is an info, the message will be printed in BLUE
*/
func (l *Logger) Info(message string) {
	l.logMessage("INFO", Blue, message)
}

/*
If there is a warning, the message will be printed in YELLOW
*/
func (l *Logger) Warn(message string) {
	l.logMessage("WARN", Yellow, message)
}

/*
If there is an error, the message will be printed in RED
*/
func (l *Logger) Error(message string) {
	l.logMessage("ERROR", Red, message)
}

func (l *Logger) Fatal(message string) {
	log.Fatalf("[FATAL] %s\n", message)
}
