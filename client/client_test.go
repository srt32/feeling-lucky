package client

import (
	"testing"
)

func TestMissingTokenError(t *testing.T) {
	emptyConfig := fakeConfig{}

	_, err := NewClient(emptyConfig)

	if err == nil {
		t.Fatalf("Missing token should return error")
	}
}

type fakeConfig struct {}

func (fake fakeConfig) Token() string {
	return ""
}
