package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidConfigFile(t *testing.T) {

	f := NewConfigFileDefault("valid-config", "yaml", "./testdata")

	assert.Equal(t, f.Config.API.Host, "0.0.0.0", "host should be 0.0.0.0 (default value)")
	assert.Equal(t, f.Config.Log.Debug, false, "log debug config should be false (default value)")
	assert.Equal(t, f.Config.API.Port, "8080", "host should be 8080 (default value)")

	err := f.LoadConfig()

	assert.Nil(t, err)
	assert.Equal(t, f.Config.API.Port, "80", "api port config should be equal")
	assert.Equal(t, f.Config.Log.Debug, true, "log debug config should be equal")
	assert.Equal(t, f.Config.API.Host, "0.0.0.0", "host should be 0.0.0.0 (default value)")
}

func TestInvalidConfigFile(t *testing.T) {

	f := NewConfigFileDefault("invalid-config", "yaml", "./testdata")

	err := f.LoadConfig()

	assert.NotNil(t, err)
}
