package gosh

import (
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseProvider Provider

func (provider DatabaseProvider) Register(application Application) Application {

	connection, _ := sqlbuilder.Open("mysql", mysql.ConnectionURL{
		Database: env("DB_NAME"),
		Host:     env("DB_HOST"),
		User:     env("DB_USER"),
		Password: env("DB_PASS"),
	})

	application = application.SetDatabaseConnection(connection)
	return application
}

func (DatabaseProvider) Boot(application Application) Application {

	// boot code here
	return application
}