package models

import (
  "github.com/jinzhu/gorm"
_ "github.com/mattn/go-sqlite3"
)

var readDB *gorm.DB

func DBInit() {
  var err error
  readDB, err = gorm.Open("sqlite3", "../models/master.db")
  if err != nil {
    panic(err.Error())
  }
}
