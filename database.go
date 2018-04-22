package gosh

import (
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Database struct {
	Adapter string
	Encoding string
	Host string
	Name string
	Username string
	Password string
}

func (database Database) Connection() (sqlbuilder.Database, error) {

	return mysql.Open(mysql.ConnectionURL{
		Database: env("DB_NAME"),
		Host:     env("DB_HOST"),
		User:     env("DB_USER"),
		Password: env("DB_PASS"),
	})
}