package models

type User struct {
  Id int
  Email string
  FirstName string
  LastName string
  PasswordHash string
  PasswordSalt string
}

