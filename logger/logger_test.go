package logger_test

import (
	"testing"

	"galaxyzeta.com/logger"
)

func TestLogger(t *testing.T) {
	log := logger.New("Test")
	log.Info("This is info")
	log.Warn("This is warn")

}
