package router

import (
	"net/http"
)

func (router *Router) init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		path := parseRoute(r.URL.Path)
		queries := parseQueries(r.URL.Query().Encode())
		ctx := &Context{W: &w, R: r, Params: make(map[string]string), Query: queries}

		length := len(*path)

		if length == 0 {
			if rt := router.routes["index"]; rt != nil {
				handleRoute(rt, ctx)
			} else {
				notFound(ctx)
			}
			return
		}

		var rt *Route

		for i, route := range *path {
			isTheLast := i == length-1

			if rt != nil {
				_rt := rt
				if rt = _rt.subroutes[route]; rt == nil {
					if rt = _rt.subroutes["key"]; rt != nil {
						ctx.Params[rt.key] = route
					} else {
						notFound(ctx)
						return
					}
				}
			} else {
				if rt = router.routes[route]; rt == nil {
					if rt = router.routes["key"]; rt != nil {
						ctx.Params[rt.key] = route
					} else {
						notFound(ctx)
						return
					}
				}
			}

			if isTheLast {
				handleRoute(rt, ctx)
				return
			}
		}

		// if params := r.URL.Query(); len(params) != 0 {
		// 	id = params["id"][0]
		// }
	})
}
