package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hurtki/routego/internal/route"
	"github.com/hurtki/routego/internal/route_set"
)

// Router is a structure to handle requests and routing them to handlers
// Router uses RouteSet to match path with handler
type Router struct {
	routeSet route_set.RouteSet
	cfg      RouterConfig
}

type RouterConfig struct {
	NotFoundHandler http.Handler
}

func defaultConfig(base RouterConfig) RouterConfig {
	cfg := RouterConfig{}
	if base.NotFoundHandler == nil {
		cfg.NotFoundHandler = http.NotFoundHandler()
	} else {
		cfg.NotFoundHandler = base.NotFoundHandler
	}

	return cfg
}

func NewRouter(cfg *RouterConfig) Router {
	if cfg == nil {
		cfg = &RouterConfig{}
	}

	return Router{
		routeSet: route_set.NewRouteSet(),
		cfg:      defaultConfig(*cfg),
	}
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	handler, parameter, found := r.routeSet.Handler(req.URL.Path, route.HttpMethod(req.Method))
	if !found {
		r.cfg.NotFoundHandler.ServeHTTP(res, req)
		return
	}

	if parameter != nil {
		ctx := context.WithValue(req.Context(), "urlParameter", parameter)
		req = req.WithContext(ctx)
	}

	handler.ServeHTTP(res, req)
}

func test(tes http.ResponseWriter, req *http.Request) {
	fmt.Println("got to handler")
	fmt.Println(req.Context().Value("urlParameter"))
}

func main() {
	router := NewRouter(nil)
	router.GetFunc("/tasks/{num}", test)
	router.PutFunc("/tasks/{num}", test)
	router.DeleteFunc("/tasks/{num}", test)
	http.ListenAndServe(":80", &router)
	time.Sleep(time.Hour)
}
