package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DSNParam string

const (
	EnvHost     = "DB_HOST"
	EnvUser     = "DB_USER"
	EnvPassword = "DB_PASSWORD"
	EnvDBName   = "DB_NAME"
	EnvPort     = "DB_PORT"
)

func NewDb() (*sql.DB, error) {
	host := os.Getenv(EnvHost)
	user := os.Getenv(EnvUser)
	password := os.Getenv(EnvPassword)
	dbname := os.Getenv(EnvDBName)
	port := os.Getenv(EnvPort)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port)
	return sql.Open("postgres", dsn)
}
