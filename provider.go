package gosh

type Provider struct {
	Application *Application
	ProviderInterface
}

type ProviderInterface interface {
	Boot()
	Register()
}