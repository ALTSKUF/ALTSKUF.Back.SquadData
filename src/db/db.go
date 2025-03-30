package db

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/models"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/schemas"
  e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/google/uuid"
  "fmt"
  "errors"
)

type Db interface {
  GetAllSquads() ([]schemas.GetSquadResponse, error)
  GetSquadById(int) schemas.GetSquadResponse
  GetSquadMembers(int) ([]uuid.UUID, error)
  Migrate()
}

type DbController struct {
  *gorm.DB
}

func Init(config *config.Config) (Db, error) {
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
    return nil, e.DbOpenError
  }

  return &DbController{db}, nil
}

func (db *DbController) Migrate() {
  db.AutoMigrate(&models.Squad{})
  db.AutoMigrate(&models.SquadMember{})
}

func (db *DbController) GetSquadById(squad_id int) schemas.GetSquadResponse  {
  var squad_info schemas.GetSquadResponse
  result := db.Model(&models.Squad{}).Where("id = ?", squad_id).First(&squad_info)
  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return schemas.GetSquadResponse{
				Error: e.DbSquadNotFoundError,
			}
    } else {
			return schemas.GetSquadResponse{
				Error: e.DbTransactionError,
			}
    }
  }

  return squad_info 
}

func (db *DbController) GetSquadMembers(squad_id int) ([]uuid.UUID, error) {
  var uuids []uuid.UUID
  result := db.Model(&models.SquadMember{}).Select("user_uuid").Where("squad_id = ?", squad_id).Find(&uuids)

  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return nil, e.DbSquadNotFoundError
    } else {
      return nil, e.DbTransactionError
    }
  }

  return uuids, nil
}

func (db *DbController) GetAllSquads() ([]schemas.GetSquadResponse, error) {
  var squads []schemas.GetSquadResponse

  result := db.Model(&models.Squad{}).Find(&squads)
  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return squads, nil
    } else {
      return nil, e.DbTransactionError
    }
  }

  return squads, nil
}
