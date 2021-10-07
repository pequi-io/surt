package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	log := New()
	assert.Equal(t, zerolog.LevelInfoValue, log.GetLevel().String(), "log level should be equal")
}

func TestNewDefaultLogger(t *testing.T) {
	log := NewDefault()
	assert.Equal(t, zerolog.LevelDebugValue, log.GetLevel().String(), "log level should be equal")
}

func TestDebugLogger(t *testing.T) {
	// set log config debug to true
	cfg.Config.Log.Debug = true

	log := New()
	assert.Equal(t, zerolog.LevelDebugValue, log.GetLevel().String(), "log level should be equal")
}
