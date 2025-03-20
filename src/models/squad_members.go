package models

import (
  "github.com/google/uuid"
)

type SquadMember struct {
  ID uint `gorm:"primaryKey"`
  SquadID uint `gorm:"not null;constraint:OnDelete:CASCADE"`
  Squad
  UserUUID uuid.UUID `gorm:"type:uuid;not null"`
}
