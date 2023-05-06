package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DSNParam string

const (
	Host     DSNParam = "127.0.0.1"
	User     DSNParam = "postgres"
	Password DSNParam = "password"
	DBName   DSNParam = "golang-nextjs-db"
	Port     DSNParam = "5432"
)

func NewDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		"127.0.0.1", User, Password, DBName, Port)

	return sql.Open("postgres", dsn)
}
