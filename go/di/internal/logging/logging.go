package logging

import (
	"fmt"
)

// L is the global instance of the logger
var L = &LoggerStdout{}

// LoggerStdout logs to std out
type LoggerStdout struct{}

// Debug logs messages at DEBUG level
func (l LoggerStdout) Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] " + message, args...)
}

// Info logs messages at INFO level
func (l LoggerStdout) Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] " + message, args...)
}

// Warn logs messages at WARN level
func (l LoggerStdout) Warn(message string, args ...interface{}) {
	fmt.Printf("[WARN] " + message, args...)
}

// Error logs messages at ERROR level
func (l LoggerStdout) Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] " + message, args...)
}