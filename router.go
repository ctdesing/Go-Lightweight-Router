package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Route struct {
	controller func(ctx *Context)
	subroutes  map[string]*Route
	method     string
	key        string
}

type Router struct {
	routes map[string]*Route
}

type Context struct {
	W      *http.ResponseWriter
	R      *http.Request
	Params map[string]string
	Query  map[string]string
}

func New() *Router {
	r := &Router{
		routes: make(map[string]*Route),
	}

	r.init()

	return r
}

func (r *Router) Route(path string, method string, controller func(ctx *Context)) {
	parsedPath := parseRoute(path)

	length := len(*parsedPath)

	fmt.Printf("FROM ROUTE: PATH: %v LENGTH: %v \n", parsedPath, length)

	var rt *Route

	if length == 0 {
		r.routes["index"] = newRoute()
		rt = r.routes["index"]
		rt.controller = controller
		rt.method = method
		return
	}

	for i, route := range *parsedPath {
		isTheLast := i == length-1

		if i == 0 && rt == nil {
			if strings.Contains(route, ":") {
				if r.routes["key"] == nil {
					r.routes["key"] = newRoute()
					r.routes["key"].key = strings.Replace(route, ":", "", -1)
				}
				rt = r.routes["key"]
			} else {
				if r.routes[route] == nil {
					r.routes[route] = newRoute()
				}
				rt = r.routes[route]
			}
		} else {
			if strings.Contains(route, ":") {
				if rt.subroutes["key"] == nil {
					rt.subroutes["key"] = newRoute()
					rt.subroutes["key"].key = strings.Replace(route, ":", "", -1)
				}
				rt = rt.subroutes["key"]
			} else {
				if rt.subroutes[route] == nil {
					rt.subroutes[route] = newRoute()
				}
				rt = rt.subroutes[route]
			}
		}

		if isTheLast {
			rt.controller = controller
			rt.method = method
			return
		}
	}
}

func ServeHTTP(port int) {
	_port := ":" + strconv.Itoa(port)

	err := http.ListenAndServe(_port, nil) // http
	if err != nil {
		fmt.Printf("http error found: %v", err)
	}
}

func ServeHTTPS(port int, certificate string, key string) {
	_port := ":" + strconv.Itoa(port)

	err := http.ListenAndServeTLS(_port, certificate, key, nil)
	if err != nil {
		fmt.Printf("https error found: %v", err)
	}
}
