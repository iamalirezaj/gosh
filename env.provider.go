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

func (EnvironmentProvider) Boot() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}
}

func (EnvironmentProvider) Register(app Application) Application {

	// your register things in application scope
	return app
}