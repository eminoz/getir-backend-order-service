package config

import "github.com/spf13/viper"

var Config *Configuration

type Configuration struct {
	MongoDb  string `mapstructure:"MONGODB_URI"`
	Port     int    `mapstructure:"PORT"`
	AppScret string `mapstructure:"APP_SECRET"`
}

func SetupConfig() (err error) {
	var configuration *Configuration
	viper.AddConfigPath("./pkg/config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&configuration)
	Config = configuration
	return
}

func GetConfig() *Configuration {
	return Config
}
