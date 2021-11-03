package router

import (
	"fmt"
	"strings"
)

func handleRoute(route *Route, ctx *Context) {
	if route.controller != nil {
		// Check if route is get or post
		if route.method == ctx.R.Method {
			route.controller(ctx)
			return
		}
	}

	notFound(ctx)
}

func notFound(ctx *Context) {
	fmt.Println("Not Found!") // REMOVE
	fmt.Fprintf(*ctx.W, "%v Not found, error 404!", ctx.R.URL.Path)
}

func parseRoute(path string) *[]string {
	_routes := strings.Split(path, "/")

	var routes []string

	for _, route := range _routes {
		if route != "" {
			routes = append(routes, route)
		}
	}

	return &routes
}

func parseQueries(queryEncoded string) map[string]string {
	queryList := strings.Split(queryEncoded, "&")

	var queryMap = make(map[string]string)

	for _, queryString := range queryList {
		if queryString != "" {
			query := strings.Split(queryString, "=")
			queryMap[query[0]] = query[1]
		}
	}

	return queryMap
}

func newRoute() *Route {
	return &Route{
		subroutes: make(map[string]*Route),
	}
}
