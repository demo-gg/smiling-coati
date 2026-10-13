package config

import (
	"os"
	"time"
)

type NotifierConfig struct {
	Workspace   string
	Channel     string
	Timeout     time.Duration
	MaxRetries  int
}

func NewConfig() *NotifierConfig {
	return &NotifierConfig{
		Workspace:  os.Getenv("SLACK_WORKSPACE"),
		Channel:    os.Getenv("SLACK_CHANNEL"),
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}
}
