package config

import "github.com/sirupsen/logrus"

func NewLogrus() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(logrus.InfoLevel)
	l.SetFormatter(&logrus.JSONFormatter{})
	return l
}
