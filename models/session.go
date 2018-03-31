package models

import (
  "time"
)

type Session struct {
  SessionKey string
  UserId int
  LoginTime time.Time
  LastActivity time.Time
}
