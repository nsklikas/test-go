package config

import (
	"time"
)

type ServerConfig struct {
	Port         int
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
	TimeoutIdle  time.Duration
}

type Config struct {
	Server ServerConfig
}

func NewConfig() Config {
	s := ServerConfig{
		Port:         8080,
		TimeoutRead:  time.Duration(30),
		TimeoutWrite: time.Duration(30),
		TimeoutIdle:  time.Duration(30),
	}
	return Config{s}
}
