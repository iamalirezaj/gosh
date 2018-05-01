package gosh

import (
	"upper.io/db.v3/mysql"
	"github.com/goshco/arrays"
	"upper.io/db.v3/lib/sqlbuilder"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseProvider Provider

func DatabaseConnection() sqlbuilder.Database {

	settings := arrays.Array(Container.Aliases["database"])

	connection , err := sqlbuilder.Open(settings.String("adapter"), mysql.ConnectionURL{
		Database: settings.String("database"),
		Host: settings.String("host"),
		User: settings.String("user"),
		Password: settings.String("password"),
	})

	if err != nil {
		panic(err.Error())
	}

	return connection
}

func (provider DatabaseProvider) Register(application Application) Application {

	return application.Alias("database", map[string] interface{} {
		"adapter": env("DB_ADAPTER"),
		"database": env("DB_NAME"),
		"host":     env("DB_HOST"),
		"user":     env("DB_USER"),
		"password": env("DB_PASS"),
	})
}

func (DatabaseProvider) Boot(application Application) Application {

	defer DatabaseConnection().Close()

	// boot code here
	return application
}