package config

import (
  "os"
  "log"
  "database/sql"
  _ "github.com/lib/pq"
)

// TODO: ensure this follows a singleton pattern
func OpenDatabase() *sql.DB {
  databaseUrl := os.Getenv("DATABASE_URL")
  db, err := sql.Open("postgres", databaseUrl)
  if err != nil {
    log.Fatal(err)
  }
  return db
}

