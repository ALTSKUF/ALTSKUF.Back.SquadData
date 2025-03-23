package app

import (
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/db"
)

type App struct {
  Db *db.DbController
}

func Init(config *config.Config) (*App, error) {
  db, err := db.Init(config)
  if err != nil {
    return nil, err
  }

  db.Migrate()

  return &App{Db: db}, nil
}
