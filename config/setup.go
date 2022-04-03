package config

import (
	"github.com/spf13/viper"
	"os"
)

func Setup() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	viper.SetDefault("APP_ENV", os.Getenv("APP_ENV"))
	viper.SetDefault("APP_URL", os.Getenv("APP_URL"))

	viper.SetDefault("DB_URI", os.Getenv("DB_URI"))
	viper.SetDefault("DB_NAME", os.Getenv("DB_NAME"))

	_ = viper.Unmarshal(&App)
	_ = viper.Unmarshal(&Cors)
	_ = viper.Unmarshal(&Database)
}
