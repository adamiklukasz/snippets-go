package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func loggerFromConfig() *zap.Logger {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 1000,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := cfg.Build()
	fmt.Printf("err=%#v\n", err)
	return logger
}

func loggerDefault() *zap.Logger {
	logger, err := zap.NewProduction()
	fmt.Printf("err=%#v\n", err)
	return logger
}

func main() {
	logger := loggerFromConfig()
	//logger := loggerDefault()
	defer logger.Sync()

	logger.Info("failed to fetch 1",
		zap.String("url", "wp.pl"),
		zap.Int("attemppt", 3),
		zap.Duration("backoff", time.Second))

	logger.DPanic("PanicD",
		zap.String("reason", "test"))

	fmt.Printf("****  EXIT ****\n")
}
