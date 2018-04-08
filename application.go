package gosh

import (
	"reflect"
	"fmt"
)

type Application struct {
	Name string
	Version float64
	Providers []map[string] ProviderInterface
	Aliases []map[string] interface{}
	RegisteredProviders []string
}

func (application Application) Make(alias string) interface{} {

	return application.FindAlias(alias)
}

func (application Application) FindAlias(alias string) interface{} {

	for index, _ := range application.Aliases {
		for name, abstract := range application.Aliases[index] {
			if name == alias {
				return abstract
			}
		}
	}

	return nil
}

func (application Application) Boot(function func()) Application {
	function()
	return application
}

func (application Application) Alias(alias string, abstract interface{}) Application {

	application.Aliases = append(application.Aliases, map[string] interface {} {
		string(alias): abstract,
	})
	return application
}

func (application Application) Register(alias string, abstract interface{}) Application {

	return application.Alias(alias, abstract)
}

func (application Application) AddProviders(providers []ProviderInterface) Application {

	for _, provider := range providers {
		application = application.AddProvider(provider)
	}

	return application
}

func (application Application) AddProvider(provider ProviderInterface) Application {

	// add provider to providers
	application.Providers = append(application.Providers, map[string] ProviderInterface {
		reflect.TypeOf(provider).Name(): provider,
	})

	return application
}

func (application Application) GetProviders() []ProviderInterface {

	BaseProviders := []ProviderInterface {
		EnvironmentProvider{},
		MysqlProvider{},
	}

	for _, providerArray := range application.Providers {
		for _, provider := range providerArray {
			BaseProviders = append(BaseProviders, provider)
		}
	}

	return BaseProviders
}

func (application Application) BootProvider(provider ProviderInterface) Application {

	provider.Boot()
	return application
}

func (application Application) RegisterProvider(provider ProviderInterface) Application {

	application = provider.Register(application)
	return application
}

func (application Application) AddToRegisteredProviders(provider ProviderInterface) Application {

	application.RegisteredProviders = append(application.RegisteredProviders, reflect.TypeOf(provider).Name())
	return application
}

func (application Application) IsProviderExists(provider ProviderInterface) bool {
	for _, registred := range application.RegisteredProviders {
		if registred == reflect.TypeOf(provider).Name() {
			return true
		}
	}

	return false
}

func (application Application) BootThenRegisterProvider(provider ProviderInterface) Application {
	application = application.RegisterProvider(provider).BootProvider(provider)
	fmt.Println(reflect.TypeOf(provider).Name() + " Boot then registered.")
	return application
}

func (application Application) RunProviderAndRegisterIt(provider ProviderInterface) Application {
	application = application.BootThenRegisterProvider(provider).AddToRegisteredProviders(provider)
	return application
}

func (application Application) RunProvider(provider ProviderInterface) Application {

	if !application.IsProviderExists(provider) {
		application = application.RunProviderAndRegisterIt(provider)
	}

	return application
}

func (application Application) Run() Application {

	for _, provider := range application.GetProviders() {
		application = application.RunProvider(provider)
	}

	return application
}