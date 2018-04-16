package config

import (
	"github.com/goproducts/errors"
	"github.com/spf13/viper"
)

var configuration *viper.Viper
var configEnvironment string

// InitializeEnvironment sets the kind of environment to use for application configuration.
func InitializeEnvironment(environment string) {
	configEnvironment = environment
}

// GetConfiguration loads and retrieves the application configuration based on the initialized environment.
func GetConfiguration() *viper.Viper {
	if configuration != nil {
		return configuration
	}

	configuration = viper.New()
	configuration.AddConfigPath("../config/")
	configuration.AddConfigPath("config/")
	configuration.SetConfigType("yaml")
	configuration.SetConfigName(configEnvironment)
	err := configuration.ReadInConfig()
	errors.HandleIfError(err)
	return configuration
}

// GetString retrieves a string configuration value.
func GetString(configValue string) string {
	return GetConfiguration().GetString(configValue)
}
