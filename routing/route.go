package routing

import (
	"gosh/response"
	"github.com/gin-gonic/gin/render"
	//"reflect"
)

type Route struct {
	Path string
	Method string
	Controller interface{}
	Action interface{}
	Name string
	RouteInterface
}

func (route Route) GetName() string {
	return route.Name
}

func (route Route) GetPath() string {
	return route.Path
}

func (route Route) GetMethod() string {
	return route.Method
}

func (route Route) GetAction() interface{} {

	action := route.Action

	//typ := reflect.TypeOf(route.Action).Kind()
	//
	//switch typ {
	//case reflect.String:
	//
	//	method := reflect.ValueOf(route.Action).String()
	//	action = reflect.ValueOf(route.Controller).
	//		MethodByName(method).Call([]reflect.Value{})[0].Interface()
	//
	//case reflect.Func:
	//
	//	method := reflect.ValueOf(route.Action)
	//	in := make([]reflect.Value, method.Type().NumIn())
	//
	//	for i := 0; i < method.Type().NumIn(); i++ {
	//		obj, _ := method.Type().In(i).MethodByName("New")
	//		in[i] = reflect.ValueOf(obj.Func)
	//	}
	//	action = method
	//}

	return action
}

func GetResponse(route RouteInterface) (int, render.Render) {

	responsing := response.Response{ StatusCode: 200, Content: route.GetAction() }
	return responsing.GetResponseWithStatusCode()
}