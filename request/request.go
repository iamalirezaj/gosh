package request

import "net/http"
import "github.com/gin-gonic/gin"

type Request http.Request

func Create() Request {
	return Request{ Method: "GET" }
}

var (
	request = Create()
	context gin.Context
)

func GetMethod() string {
	return request.Method
}

func All() interface{} {

	req, _ := http.NewRequest("POST", "/", nil)

	return req.URL
}