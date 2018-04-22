package routing

type DELETE Route

func (route DELETE) GetName() string {
	return route.Name
}

func (route DELETE) GetMethod() string {
	return "DELETE"
}

func (route DELETE) GetPath() string {
	return route.Path
}

func (route DELETE) GetAction() interface{} {
	return route.Action
}
