package db

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/models"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/dto"
  e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/google/uuid"
  "fmt"
  "errors"
)

type DbController struct {
  *gorm.DB
}

func Init(config *config.Config) (*DbController, error) {
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

  return &DbController{db}, nil
}

func (db *DbController) GetSquadInfo(squad_id int) (*dto.SquadInfo, error) {
  var squad_info dto.SquadInfo
  result := db.Model(&models.Squad{}).Where("id = ?", squad_id).First(&squad_info)
  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return nil, nil
    } else {
      return nil, e.DbTransactionError
    }
  }

  return &squad_info, nil
}

func (db *DbController) GetSquadMembers(squad_id int) ([]uuid.UUID, error) {
  var uuids []uuid.UUID
  result := db.Model(&models.SquadMember{}).Select("user_uuid").Where("squad_id = ?", squad_id).Find(&uuids)

  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return uuids, nil
    } else {
      return nil, e.DbTransactionError
    }
  }

  return uuids, nil
}
