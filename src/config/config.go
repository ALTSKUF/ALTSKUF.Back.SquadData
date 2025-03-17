package config

import (
  "github.com/spf13/viper"
)

type Config struct {
  AppAddress string
}

func setDefaults() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")
  viper.AutomaticEnv()

  viper.SetDefault("app.address", ":8000")
}

func Default() *Config {
  setDefaults()

  viper.ReadInConfig()

  appAddress := viper.GetString("app.address")

  return &Config {
    AppAddress: appAddress,
 }
}
