package gosh

import (
	"github.com/gin-gonic/gin"
	"gosh/routing"
)

type RouterProvider Provider

func (provider RouterProvider) Register(app Application) Application {

	var router = gin.Default()

	for _, route := range app.Routes {
		if route.GetMethod() == "GROUP" {
			router = provider.AddGroupRoute(router, route.GetPath(), route.GetRoutes())
		} else {
			provider.AddRoute(router, route)
		}
	}

	return app.SetRouter(router)
}

func (provider RouterProvider) AddGroupRoute(router *gin.Engine, path string, routes []routing.RouteInterface) *gin.Engine {

	for _, route := range routes {
		provider.AddRoute(router.Group(path), route)
	}

	return router
}

func (RouterProvider) AddRoute(router gin.IRouter, route routing.RouteInterface) gin.IRoutes {

	action := func(context *gin.Context) {
		context.Render(routing.GetResponse(route))
	}

	return router.Handle(route.GetMethod(),route.GetPath(),action)
}

func (p RouterProvider) Boot(application Application) Application {

	application.GetRouter().Run("127.0.0.1:8000")

	return application
}