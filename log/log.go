package log

import (
	"go.uber.org/zap"
)

type Config struct {
}

type Logger struct {
	*zap.Logger
}

var logger Logger

func DefaultConfig() (Config, error) {
	return Config{}, nil
}

func Init(conf Config) error {
	return nil
}