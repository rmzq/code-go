package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("this is info log")
	logrus.Debug("this is debug log")
	logrus.Error("this is error log")
}
