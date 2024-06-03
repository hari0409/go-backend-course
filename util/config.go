package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	// Path
	viper.AddConfigPath(path)
	// File Name
	viper.SetConfigName("app")
	// File extension type
	viper.SetConfigType("env")
	// Override the values in ENV if found than using the file
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	// Unmarshall the value into the config object
	err = viper.Unmarshal(&config)
	// Empty return because of named return
	// Thus it will automatically return the values
	return
}
