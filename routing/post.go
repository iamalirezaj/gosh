package routing

type POST Route

func (route POST) GetName() string {
	return route.Name
}

func (route POST) GetMethod() string {
	return "POST"
}

func (route POST) GetPath() string {
	return route.Path
}

func (route POST) GetAction() interface{} {
	return route.Action
}
