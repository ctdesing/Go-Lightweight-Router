# Go-Lightweight-Router
Lightweight Router for Golang using net/http standard library with custom route parsing, handler and context.



```go
package main

import router "github.com/ctdesing/Go-Lightweight-Router"

func main() {
	r := router.New()

	r.GET("/", func(ctx *router.Context) {
		fmt.Fprintf(*ctx.W, "Welcome to index")
	})

	r.GET("/:id", func(ctx *router.Context) {
		fmt.Fprintf(*ctx.W, "Welcome to %v", ctx.Params["id"])
	})

	r.GET("/:id/register/:product", func(ctx *router.Context) {
		fmt.Fprintf(
			*ctx.W,
			"We're going to register %v to %v, are you sure?", ctx.Params["product"], ctx.Params["id"],
		)
	})

	r.GET("/:id/register", func(ctx *router.Context) {
		fmt.Fprintf(*ctx.W, "Welcome to %v register section", ctx.Params["id"])
	})

	r.Route("/user", http.MethodGet, func(ctx *router.Context) {
		fmt.Fprintf(*ctx.W, "Welcome to user %v", ctx.Query["user"])
	})

	r.POST("/user/app", func(ctx *router.Context) {
		fmt.Fprintf(*ctx.W, "Welcome to user app")
	})

	router.ServeHTTP(3000)
}

```