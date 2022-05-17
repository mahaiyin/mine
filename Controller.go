package mine

import (
	"mine/controller"
	"net/http"
)

var Controller = map[string]func(http.ResponseWriter, *http.Request){
	"controller.Home":      controller.Home,
	"controller.Status404": controller.Status404,
}
