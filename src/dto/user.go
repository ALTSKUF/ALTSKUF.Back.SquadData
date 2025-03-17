package models

import (
  "github.com/google/uuid"
)

type User struct {
  UUID uuid.UUID 
  FullName string 
  Group string 
}
