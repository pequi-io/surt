package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	envPrefix       = strings.ToUpper(surtEnvPrefix)
	envName         = envPrefix + "_NAME"
	envPath         = envPrefix + "_PATH"
	validFileName   = "valid-config"
	invalidFileName = "invalid-config"
	testdataPath    = "./testdata"
	// Should reflect the values in `default:"value"`
	defaultFileTest = File{
		Name:      "config",
		Extension: "yaml",
		Path:      "$HOME/.surt",
		Config: Config{
			API: API{
				Port: "8080",
			},
			Log: Log{
				Pretty: false,
				Debug:  false,
			},
		},
	}
)

func TestDefaultFile(t *testing.T) {

	// make sure env vars are unset
	cleanup()

	// Use default
	f, err := Default()

	// Assert default values for File
	assert.Nil(t, err)
	assert.Equal(t, defaultFileTest.Name, f.Name, "file name should be equal (default value)")
	assert.Equal(t, defaultFileTest.Extension, f.Extension, "file extension should be equal (default value)")
	assert.Equal(t, defaultFileTest.Path, f.Path, "file path should be equal (default value)")

}

func TestDefaultFileWithEnvVars(t *testing.T) {

	setup(validFileName)

	// Use default
	f, err := Default()

	// Assert default values for File
	assert.Nil(t, err)
	assert.Equal(t, validFileName, f.Name, "file name should be equal")
	assert.Equal(t, defaultFileTest.Extension, f.Extension, "file extension should be equal (default value)")
	assert.Equal(t, testdataPath, f.Path, "file path should be equal")

}

func TestValidConfigFile(t *testing.T) {

	// Create env vars with valid config file
	setup(validFileName)

	f, err := Default()

	// Assert default values
	assert.Nil(t, err)
	assert.Equal(t, defaultFileTest.Config.Log.Debug, f.Config.Log.Debug, "log debug config should be equal (default value)")
	assert.Equal(t, defaultFileTest.Config.Log.Pretty, f.Config.Log.Pretty, "log pretty config should be equal (default value)")
	assert.Equal(t, defaultFileTest.Config.API.Port, f.Config.API.Port, "host should be equal (default value)")

	err = f.LoadConfig()

	// Assert default and new values from config file
	assert.Nil(t, err)
	assert.Equal(t, true, f.Config.Log.Debug, "log debug config should be equal")
	assert.Equal(t, defaultFileTest.Config.Log.Pretty, f.Config.Log.Pretty, "log pretty config should be false (default value)")
	assert.Equal(t, "80", f.Config.API.Port, "api port config should be equal")

}

func TestInvalidConfigFile(t *testing.T) {

	// Create env vars with invalid config file
	setup(invalidFileName)

	f, err := Default()

	assert.Nil(t, err)

	err = f.LoadConfig()

	// load from file should faid
	assert.NotNil(t, err)
}

func setup(fileName string) {

	cleanup()

	// set env vars for test
	os.Setenv(envName, fileName)
	os.Setenv(envPath, testdataPath)

}

func cleanup() {
	// Unset env vars for test
	os.Unsetenv(envName)
	os.Unsetenv(envPath)
}
