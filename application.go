package gosh

import (
	"reflect"
	"gosh/routing"
	"gosh/request"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Application struct {
	Name string
	Version float64
	Mode string
	Providers []ProviderInterface
	Routes []routing.RouteInterface
	Request request.Request
	Router *gin.Engine
	Connection sqlbuilder.Database
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

func (application Application) GetProviders() []ProviderInterface {

	BaseProviders := []ProviderInterface {}

	BeforProvider := []ProviderInterface {
		EnvironmentProvider{},
		MysqlProvider{},
	}

	AfterProviders := []ProviderInterface {
		RouterProvider{},
	}

	BaseProviders = append(BeforProvider, application.Providers... )
	BaseProviders = append(BaseProviders, AfterProviders...)

	return BaseProviders
}

func (application Application) BootProvider(provider ProviderInterface) Application {

	provider.Boot(application)

	if application.Mode == "debug" {
		color.Green(reflect.TypeOf(provider).String() + " booted.")
	}

	return application
}

func (application Application) RegisterProvider(provider ProviderInterface) Application {

	application = provider.Register(application)

	if application.Mode == "debug" {
		color.Green(reflect.TypeOf(provider).String() + " registered.")
	}

	return application
}

func (application Application) RunProviders() Application {

	providers := application.GetProviders()

	application = application.RegisterProviders(providers).BootProviders(providers)

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

	application = application.RunProviders()
	return application
}