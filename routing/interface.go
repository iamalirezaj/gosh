package routing

type RouteInterface interface {
	GetName() string
	GetPath() string
	GetMethod() string
	GetAction() interface{}
	GetRoutes() []RouteInterface
}