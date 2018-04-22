package gosh

type Provider struct {
	ProviderInterface
	Application Application
}

type ProviderInterface interface {
	Boot(app Application) Application
	Register(app Application) Application
}