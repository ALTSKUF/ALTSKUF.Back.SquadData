package models

import (
  "github.com/google/uuid"
)

type SquadMember struct {
  ID uint `gorm:"primaryKey"`
  SquadID uint `gorm:"not null"`
  Squad
  UserUUID uuid.UUID `gorm:"type:uuid;not null"`
  Role string `gorm:"not null"`
}
