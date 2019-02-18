package gosh

import "os"

func Env(value string) string {
	return os.Getenv(value)
}

func App() *Application {
	return Container
}

func GetConfig(key string) interface{} {
	return Container.Configs[key]
}

func SetConfig(key string, value interface{}) {
	Container.Configs[key] = value
}
