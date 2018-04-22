package routing

type PUT Route

func (route PUT) GetName() string {
	return route.Name
}

func (route PUT) GetMethod() string {
	return "PUT"
}

func (route PUT) GetPath() string {
	return route.Path
}

func (route PUT) GetAction() interface{} {
	return route.Action
}
