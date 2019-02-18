package gosh

import (
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseProvider Provider

func (provider DatabaseProvider) Register() {

	SetConfig("database", map[string]interface{}{
		"adapter":  Env("DB_ADAPTER"),
		"database": Env("DB_NAME"),
		"host":     Env("DB_HOST"),
		"user":     Env("DB_USER"),
		"password": Env("DB_PASS"),
	})
}