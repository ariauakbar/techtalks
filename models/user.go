package models

import (
  "github.com/jinzhu/gorm"
  //"log"
)

type User struct {
  gorm.Model
  FullName string
  Email string `json:"email"`
  GoogleIDToken string
  ImageURL string
}

// func CreateUser(name string) *User {
//   u := &User{name}
//   return u
// }

func CheckIfExisting(u *User) bool {
  result := readDB.Debug().First(&u)
  if result.Error == gorm.ErrRecordNotFound {
    return false
  }
  return true
}

func Register(u *User) {
  readDB.Debug().Create(&u)
}


func CreateUserTable() {
  readDB.Debug().AutoMigrate(&User{})
}
