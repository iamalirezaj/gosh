package routing

type GROUP struct {
	Path string
	Routes []RouteInterface
	RouteInterface
}

func (route GROUP) GetMethod() string {
	return "GROUP"
}

func (route GROUP) GetRoutes() []RouteInterface {
	return route.Routes
}

func (route GROUP) GetPath() string {
	return route.Path
}