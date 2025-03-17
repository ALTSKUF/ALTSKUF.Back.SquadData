package config

import (
  "github.com/spf13/viper"
)

type Config struct {
  AppAddress string
  AppProfile string
}

func setDefaults() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")

  viper.SetDefault("address", ":8000")
  viper.SetDefault("profile", "debug")

  viper.AutomaticEnv()
  viper.SetEnvPrefix("squad")
}

func Default() *Config {
  setDefaults()

  viper.ReadInConfig()

  appAddress := viper.GetString("address")
  appProfile := viper.GetString("profile")

  return &Config {
    AppAddress: appAddress,
    AppProfile: appProfile,
  }
}
