package models

type User struct {
  Name string
}

func CreateUser(name string) *User {
  u := &User{name}
  return u
}
