package main

import "galaxyzeta.com/logger"

func main() {
	logger := logger.New("TestLogger")
	logger.Info("Hello")
	logger.Infof("Hello %s", "world")
	logger.Debug("Hello")
	logger.Debugf("Hello %s", "world")
	logger.Warn("Hello")
	logger.Warnf("Hello %s", "world")
	logger.Error("Hello")
	logger.Errorf("Hello %s", "world")
	logger.Fatal("Hello")
	logger.Fatalf("Hello %s", "world")
}
