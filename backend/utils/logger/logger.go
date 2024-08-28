package logger

import "github.com/sirupsen/logrus"

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func Info(msg string, args ...any) {
	logrus.Infof(msg, args...)
}

func Warn(msg string, args ...any) {
	logrus.Warnf(msg, args...)
}

func Error(msg string, args ...any) {
	logrus.Errorf(msg, args...)
}
