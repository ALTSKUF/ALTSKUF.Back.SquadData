package db

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

type Db struct {
  *gorm.DB
}

func InitDb(config *config.Config) (*Db, error) {
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    config.DbHost,
    config.DbUser,
    config.DbPassword,
    config.DbName,
    config.DbPort,
    config.DbSSLMode,
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, apperror.DbOpenError
  }

  return &Db{db}, nil
}
