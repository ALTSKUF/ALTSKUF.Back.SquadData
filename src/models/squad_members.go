package models

import (
  "github.com/google/uuid"
)

type SquadMember struct {
  ID uint `gorm:"primaryKey"`
	UserUUID uuid.UUID `gorm:"type:uuid;not null;default:gen_random_uuid()"`
  SquadID uint `gorm:"not null;constraint:OnDelete:CASCADE"`
  Squad Squad
}
