package config

import (
  "github.com/spf13/viper"
  "strings"
)

type Config struct {
  AppAddress string
  AppProfile string
  DbHost string
  DbUser string
  DbPassword string
  DbName string
  DbPort string
  DbSSLMode string
  RMQUser string
  RMQPassword string
  RMQHost string
  RMQPort string
}

func setDefaults() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")

  viper.SetDefault("app.address", ":8000")
  viper.SetDefault("app.profile", "debug")

  viper.SetDefault("db.port", "5432")
  viper.SetDefault("db.host", "db")
  viper.SetDefault("db.user", "postgres")
  viper.SetDefault("db.name", "postgres")
  viper.SetDefault("db.sslmode", "disable")

  viper.SetDefault("rmq.port", "5672")

  viper.AutomaticEnv()
  viper.SetEnvPrefix("squad")
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func Default() *Config {
  setDefaults()

  viper.ReadInConfig()

  appAddress := viper.GetString("app.address")
  appProfile := viper.GetString("app.profile")

  dbPort := viper.GetString("db.port")
  dbHost := viper.GetString("db.host")
  dbPassword := viper.GetString("db.password")
  dbName := viper.GetString("db.name")
  dbUser := viper.GetString("db.user")
  dbSSLMode := viper.GetString("db.sslmode")

  rmqUser := viper.GetString("rmq.user")
  rmqPassword := viper.GetString("rmq.password")
  rmqHost := viper.GetString("rmq.host")
  rmqPort := viper.GetString("rmq.port")

  return &Config {
    AppAddress: appAddress,
    AppProfile: appProfile,
    DbHost: dbHost,
    DbUser: dbUser,
    DbPassword: dbPassword,
    DbName: dbName,
    DbPort: dbPort,
    DbSSLMode: dbSSLMode,
    RMQUser: rmqUser,
    RMQPassword: rmqPassword,
    RMQHost: rmqHost,
    RMQPort: rmqPort,
  }
}
