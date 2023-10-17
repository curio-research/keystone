package logging

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	var err error

	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Log() *zap.SugaredLogger {
	return logger.Sugar()
}

func BasicLog() *zap.Logger {
	return logger
}
