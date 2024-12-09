package logger

import "testing"

func TestLogLevels(t *testing.T) {
	l := NewLogger("DEBUG")

	l.Debug("Debug message")
	l.Info("Info message")
	l.Warn("Warn message")
	l.Error("Error message")
	l.Fatal("Fatal message")
}
