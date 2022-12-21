package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(Init)

func Init() (logger *logrus.Logger) {
	logFile := "../log.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		return
	}
	defer f.Close()

	logger = &logrus.Logger{
		Out:       f,
		Formatter: &logrus.JSONFormatter{},
		Level:     0,
	}

	return
}
