package config

import (
	"os"
)

type Config interface {
	Token() string
}

type UserConfig struct {}

func (c UserConfig) Token() string {
	return os.Getenv("GITHUB_PERSONAL_TOKEN")
}
