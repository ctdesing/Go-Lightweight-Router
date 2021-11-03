package router

import "net/http"

func (r *Router) GET(path string, controller func(ctx *Context)) {
	r.Route(path, http.MethodGet, controller)
}

func (r *Router) POST(path string, controller func(ctx *Context)) {
	r.Route(path, http.MethodPost, controller)
}

func (r *Router) PUT(path string, controller func(ctx *Context)) {
	r.Route(path, http.MethodPut, controller)
}

func (r *Router) DELETE(path string, controller func(ctx *Context)) {
	r.Route(path, http.MethodDelete, controller)
}
