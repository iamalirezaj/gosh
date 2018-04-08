package gosh

import (
	"log"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
)

type MysqlProvider Provider

func (MysqlProvider) Boot() {}

func (provider MysqlProvider) Register(app Application) Application {

	var Database sqlbuilder.Database
	var DBError error

	Database, DBError = mysql.Open(mysql.ConnectionURL{
		Database: env("DB_NAME"),
		Host:     env("DB_HOST"),
		User:     env("DB_USER"),
		Password: env("DB_PASS"),
	})

	if DBError != nil {
		log.Fatal("Mysql connection Error: ", DBError)
	}

	Database.SetLogging(false)

	defer Database.Close()

	return app.Register("database", Database)
}