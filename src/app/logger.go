package app

import "github.com/sirupsen/logrus"

type Logger struct {
	logger  *logrus.Logger
	appInfo logrus.Fields
}

func InitLogger() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{
		logger: logger,
		appInfo: logrus.Fields{
			"service": "boiler-plate",
			"module":  "ping",
			"version": "0.0.1-beta",
		},
	}
}

func (l *Logger) LogSuccess(category string, method string, message string) {
	l.appInfo["category"] = category
	l.appInfo["method"] = method
	l.logger.WithFields(l.appInfo).Info(message)
}
func (l *Logger) LogSuccessWithAdditionalInfo(category string, method string, message string, additionalInfo map[string]interface{}) {
	l.appInfo["category"] = category
	l.appInfo["method"] = method
	l.appInfo["additional_info"] = additionalInfo
	l.logger.WithFields(l.appInfo).Info(message)
}
func (l *Logger) LogError(err error) {
	l.logger.WithFields(l.appInfo).Error(err)
}
