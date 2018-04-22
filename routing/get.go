package routing

type GET Route

func (route GET) GetName() string {
	return route.Name
}

func (route GET) GetMethod() string {
	return "GET"
}

func (route GET) GetPath() string {
	return route.Path
}

func (route GET) GetAction() interface{} {
	return route.Action
}
