package models

import (
  "github.com/jinzhu/gorm"
_ "github.com/mattn/go-sqlite3"
)

func DB() (*gorm.DB, error) {
  db, err := gorm.Open("sqlite3", "../models/master.db")
  return db, err
}
