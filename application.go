package gosh

import (
	"reflect"
	"io/ioutil"
	"gosh/routing"
	"gosh/request"
	"gopkg.in/yaml.v2"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"upper.io/db.v3/lib/sqlbuilder"
)

var Container = Application{ Name: "Gosh", Version: "0.0.1" }

type Application struct {
	Name string
	Version string
	Environment string
	Providers []ProviderInterface
	Routes []routing.RouteInterface
	Request request.Request
	Router *gin.Engine
	Aliases []Alias
	Connection sqlbuilder.Database
}

func(application Application) LoadConfigs(filename string) Application {
	yamlFile, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &application)
	return application
}

func (application Application) Alias(alias Alias) Application {
	application.Aliases = append(application.Aliases, alias)
	return application
}

func (application Application) Make(action interface{}) Application {

	return application
}

func (application Application) SetRouter(router *gin.Engine) Application {

	application.Router = router
	return application
}

func (application Application) SetDatabaseConnection(connection sqlbuilder.Database) Application {

	application.Connection = connection
	return application
}

func (application Application) GetRouter() *gin.Engine {
	return application.Router
}

func (application Application) SetRoutes(routes []routing.RouteInterface) Application {

	application.Routes = append(application.Routes, routes...)
	return application
}

func (application Application) AddProviders(providers []ProviderInterface) Application {

	application.Providers = append(application.Providers, providers...)
	return application
}

func (application Application) SetProviders() Application {

	BeforProvider := []ProviderInterface {
		EnvironmentProvider{},
		DatabaseProvider{},
	}

	AfterProviders := []ProviderInterface {
		RouterProvider{},
	}

	application.Providers = append(BeforProvider, application.Providers... )
	application.Providers = append(application.Providers, AfterProviders...)

	return application
}

func (application Application) BootProvider(provider ProviderInterface) Application {

	application = provider.Boot(application)
	Container = application

	if application.Environment == "debug" {
		color.Green(reflect.TypeOf(provider).String() + " booted.")
	}

	return application
}

func (application Application) RegisterProvider(provider ProviderInterface) Application {

	application = provider.Register(application)
	Container = application

	if application.Environment == "debug" {
		color.Green(reflect.TypeOf(provider).String() + " registered.")
	}

	return application
}

func (application Application) RunProviders() Application {

	application = application.SetProviders()
	application = application.RegisterProviders(application.Providers).BootProviders(application.Providers)

	return application
}

func (application Application) BootProviders(providers []ProviderInterface) Application {

	for _, provider := range providers {
		application = application.BootProvider(provider)
	}
	return application
}

func (application Application) RegisterProviders(providers []ProviderInterface) Application {

	for _, provider := range providers {
		application = application.RegisterProvider(provider)
	}
	return application
}

func (application Application) Run() Application {

	return application.RunProviders()
}