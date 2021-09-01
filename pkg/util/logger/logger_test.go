package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLogger(t *testing.T) {
	log := New(false, false)
	assert.Equal(t, zerolog.LevelInfoValue, log.GetLevel().String(), "log level should be equal")
}

func TestDebugLogger(t *testing.T) {
	log := New(true, false)
	assert.Equal(t, zerolog.LevelDebugValue, log.GetLevel().String(), "log level should be equal")
}
