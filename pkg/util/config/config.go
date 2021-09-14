package config

import (
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

// File defines some properties of config file
type File struct {
	Name      string `default:"config"`
	Extension string `default:"yaml"`
	Path      string `default:"$HOME/.surt"`
	Config    Config
}

// Global config
type Config struct {
	API       API
	Log       Log
	Antivirus Antivirus
}

// API config
type API struct {
	Port string `default:"8080"`
}

// Global Log config
type Log struct {
	Debug  bool `default:"false"`
	Pretty bool `default:"false"`
}

type Antivirus struct {
	Engine  string `default:"clamav"`
	Address string `default:"127.0.0.1:3310"`
	Network string `default:"tcp"`
}

// Environment variable prefix
var surtEnvPrefix = "surt_cfg"

// Return default File with default config
func Default() (*File, error) {

	// New file using default values
	f := new(File)
	defaults.SetDefaults(f)

	v := viper.New()

	// load environment variables
	// finding for env vars: SURT_CFG_NAME, SURT_CFG_PATH and SURT_CFG_EXTENSION
	v.AutomaticEnv()
	v.SetEnvPrefix(surtEnvPrefix)

	// We must specify the env vars that we want to use in unmarshals using `BindEnv` based on the following issue:
	// https://github.com/spf13/viper/issues/188
	v.BindEnv("name")
	v.BindEnv("path")

	// viper.Unmarshal unmarshals into a f struct
	err := v.Unmarshal(&f)
	if err != nil {
		return f, err
	}

	return f, nil
}

func (f *File) LoadConfig() error {

	v := viper.New()

	v.SetConfigName(f.Name)
	v.SetConfigType(f.Extension)
	v.AddConfigPath(f.Path + "/")

	// load configuration file
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// load environment variables
	v.AutomaticEnv()
	v.SetEnvPrefix(surtEnvPrefix)

	// viper.Unmarshal unmarshals the config into a &f.Config struct
	err := v.Unmarshal(&f.Config)

	if err != nil {
		return err
	}
	return nil
}
