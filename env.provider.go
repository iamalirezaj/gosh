package gosh

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func env(value string) string {

	return os.Getenv(value)
}

type EnvironmentProvider Provider

func (EnvironmentProvider) Register(application Application) Application {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	application.Mode = env("ENV")
	if application.Mode == "" { application.Mode = "debug" }

	// your register things in application scope
	return application
}

func (EnvironmentProvider) Boot(application Application) Application {

	return application
}