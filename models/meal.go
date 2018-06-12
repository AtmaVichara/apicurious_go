package models

import (
  "time"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Meal struct {
  gorm.Model

  Name string `gorm: "type:varchar(100)" json: "name"`
}