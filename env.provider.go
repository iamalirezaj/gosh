package gosh

import (
	"github.com/pkg/errors"
	"github.com/joho/godotenv"
)

type EnvironmentProvider Provider

func (provider EnvironmentProvider) Register() {

	if err := godotenv.Load(); err != nil {
		errors.Errorf("%v", err)
	}

	provider.Application.Environment = Env("ENV")

	if provider.Application.Environment == "" {
		provider.Application.Environment = "debug"
	}
}