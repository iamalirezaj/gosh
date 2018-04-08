package gosh

import (
	"os"
	"github.com/gin-gonic/gin"
)

type BaseRouterProvider struct {
	ProviderInterface
}

func (BaseRouterProvider) Boot(app Application) Application {
	gin.SetMode(os.Getenv("GIN_MODE"))
	app.Make("router")
	return app
}

func (provider BaseRouterProvider) Register(app Application) Application {

	app = app.Register("router", func() interface{} {

		return gin.Default()
	})

	return app
}
