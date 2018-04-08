package gosh

type ProviderInterface interface {
	Boot()
	Register(application Application) Application
}

type Provider struct {
	ProviderInterface
}