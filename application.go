package gosh

import (
	"reflect"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Application struct {
	Name string
	Version string
	Environment string
	Providers []ProviderInterface
	Router *gin.Engine
	Configs map[string] interface{}
	Connection sqlbuilder.Database
}

var Container *Application

func NewApplication(name string) *Application {

	Container = &Application{
		Name: "Gosh",
		Version: "0.0.1",
		Configs: map[string] interface{} {},
	}
	return Container
}

func(application Application) LoadConfigs(filename string) Application {
	yamlFile, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &application)
	return application
}

func (application *Application) Config(key string, value interface{}) {
	application.Configs[key] = value
}

func (application *Application) SetRouter(router *gin.Engine) {
	application.Router = router
}

func (application *Application) SetDatabaseConnection(connection sqlbuilder.Database) {
	application.Connection = connection
}

func (application *Application) GetRouter() *gin.Engine {
	return application.Router
}

func (application *Application) AddProvider(provider ProviderInterface) {
	application.Providers = append(application.Providers, provider)
}

func (application *Application) AddProviders(providers []ProviderInterface) {
	application.Providers = append(application.Providers, providers...)
}

func (application *Application) SetProviders() {

	BaseProviders := []ProviderInterface {
		&EnvironmentProvider{},
		&DatabaseProvider{},
	}

	application.Providers = append(BaseProviders, application.Providers... )
}

func (application *Application) BootProvider(provider ProviderInterface) {

	provider.Boot()

	if application.Environment == "debug" {
		color.Green(reflect.TypeOf(provider).String() + " booted.")
	}
}

func (application *Application) RegisterProvider(provider ProviderInterface) {

	p := reflect.ValueOf(provider).Elem()

	d := p.FieldByName("Application")
	d.Set(reflect.ValueOf(application))

	p.MethodByName("Register").Call([] reflect.Value{})

	if application.Environment == "debug" {
		color.Green(reflect.TypeOf(provider).String() + " registered.")
	}
}

func (application *Application) RunProviders() {

	application.SetProviders()
	application.RegisterProviders()
	//application.BootProviders(application.Providers)
}

func (application *Application) BootProviders(providers []ProviderInterface) {

	for _, provider := range providers {
		application.BootProvider(provider)
	}
}

func (application *Application) RegisterProviders() {

	for _, provider := range application.Providers {
		application.RegisterProvider(provider)
	}
}

func (application *Application) Run() {
	application.RunProviders()
}