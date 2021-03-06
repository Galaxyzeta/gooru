package logger

// Logger uses Sprintf to log, so it is NOT efficient enough.
// Please use ZAP as logger framework if you are demanding at logging efficiency !

import (
	"fmt"
	"time"
)

const (
	noColor   = 0
	bgBlack   = 30
	bgRed     = 31
	bgGreen   = 32
	bgYellow  = 33
	bgBlue    = 34
	bgMagenta = 35
	bgCyan    = 36
	bgWhite   = 37
	fgBlack   = 40
	fgRed     = 41
	fgGreen   = 42
	fgYellow  = 43
	fgBlue    = 44
	fgMagenta = 45
	fgCyan    = 46
	fgWhite   = 47
)

// Logger is a representation of logging util.
type Logger struct {
	enable bool
	name   string
}

// New brings you a new logger.
func New(name string) *Logger {
	return &Logger{enable: true, name: name}
}

// Enable the logger.
func (logger *Logger) Enable() {
	logger.enable = true
}

// Disable the logger.
func (logger *Logger) Disable() {
	logger.enable = false
}

// Info gives some hint.
func (logger *Logger) Info(msg string) {
	logger.doLog("[INFO]", msg, bgBlue, fgBlack)
}

// Infof gives some hint with format.
func (logger *Logger) Infof(format string, arg ...interface{}) {
	logger.Info(fmt.Sprintf(format, arg...))
}

// Debug gives some debug info.
func (logger *Logger) Debug(msg string) {
	logger.doLog("[DEBUG]", msg, bgWhite, fgBlack)
}

// Debugf gives some debug info with format.
func (logger *Logger) Debugf(format string, arg ...interface{}) {
	logger.Debug(fmt.Sprintf(format, arg...))
}

// Warn gives some warning info.
func (logger *Logger) Warn(msg string) {
	logger.doLog("[WARN]", msg, bgYellow, fgBlack)
}

// Warnf gives some warning info with format.
func (logger *Logger) Warnf(format string, arg ...interface{}) {
	logger.Warn(fmt.Sprintf(format, arg...))
}

// Error gives some error info.
func (logger *Logger) Error(msg string) {
	logger.doLog("[ERROR]", msg, bgRed, fgBlack)
}

// Errorf gives some error info with format.
func (logger *Logger) Errorf(format string, arg ...interface{}) {
	logger.Error(fmt.Sprintf(format, arg...))
}

// Fatal gives some critical info.
func (logger *Logger) Fatal(msg string) {
	logger.doLog("[FATAL]", msg, bgMagenta, fgBlack)
}

// Fatalf gives some critical info with format.
func (logger *Logger) Fatalf(format string, arg ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, arg...))
}

func (logger *Logger) doLog(label string, msg string, bgcolor int, fgcolor int) {
	if logger.enable == false {
		return
	}
	fmt.Printf("%c[0;%dm%s\t%d | %s:\t%s%c[0m\n", 0x1B, bgcolor, label, time.Now().Unix(), logger.name, msg, 0x1B)
}
