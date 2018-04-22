package response

import (
	"reflect"
	"github.com/gin-gonic/gin/render"
)

type Response struct {
	StatusCode int
	Content interface{}
}

func (r Response) GetResponseWithStatusCode() (int, render.Render) {

	Type := reflect.TypeOf(r.Content).Kind();
	switch Type {

	case reflect.String:
		return r.GetJsonRender(r.StatusCode, r.Content)

	case reflect.Func:
		return r.GetFunctionRender(r.StatusCode, r.Content )

	case reflect.Slice,
	reflect.Map,
	reflect.Array:
		return r.GetJsonSliceRender(r.StatusCode, r.Content)

	default:
		return r.GetInterfaceRender(r.StatusCode, r.Content)
	}
}

func (r Response) GetInterfaceRender(code int, data interface{}) (int, render.Render) {

	return code, render.JSON{Data: data}
}

func (r Response) GetFunctionRender(code int, data interface{}) (int, render.Render) {

	return code, render.JSON{Data: data}
}

func (r Response) GetJsonRender(code int, data interface{}) (int, render.Render) {

	return code, render.JSON{Data: data}
}

func (r Response) GetJsonSliceRender(code int, data interface{}) (int, render.Render) {

	return code, render.JSON{Data: data}
}