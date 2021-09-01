package config

import (
	"github.com/spf13/viper"
)

// File defines some properties of config file
type File struct {
	Name      string
	Extension string
	Path      string
	Config    Config
}

// Global config
type Config struct {
	API API
	Log Log
}

// API config
type API struct {
	Host string
	Port string
}

// Global Log config
type Log struct {
	Debug  bool
	Pretty bool
}

// NewConfigFileDefault creates new File struct with Config default values
// viper doesn't support default values in struct: https://github.com/spf13/viper/issues/295
func NewConfigFileDefault(name string, extension string, path string) File {
	return File{
		Name:      name,
		Extension: extension,
		Path:      path,
		Config: Config{
			API: API{
				Host: "0.0.0.0",
				Port: "8080",
			},
			Log: Log{
				Debug:  false, // default: info
				Pretty: false, // default: json
			},
		},
	}
}

func (f *File) LoadConfig() error {

	viper.SetConfigName(f.Name)
	viper.SetConfigType(f.Extension)
	viper.AddConfigPath(f.Path + "/")

	// load configuration file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// viper.Unmarshal unmarshals the config into a &f.Config Struct
	err := viper.Unmarshal(&f.Config)
	if err != nil {
		return err
	}
	return nil
}
