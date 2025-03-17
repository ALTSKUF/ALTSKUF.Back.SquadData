package models

import (
  "gorm.io/gorm"
)

type Squad struct {
  gorm.Model 
  Name string `gorm:"not null;unique;size:50"`
  Description string
}
