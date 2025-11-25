package routego

import (
	"context"
	"net/http"

	"github.com/hurtki/routego/route"
)

// Router is a structure to handle requests and routing them to handlers
// Router uses RouteSet to match path with handler
type Router struct {
	routeSet RouteSet
}

func NewRouter(routeSet RouteSet) Router {

	router := Router{
		routeSet: routeSet,
	}

	return router
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	handler, parameter, err := r.routeSet.Handler(req.URL.Path)
	if err == route.ErrNotFound {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.NotFound(res, req)
	}

	if parameter != nil {
		ctx := context.WithValue(req.Context(), "urlParameter", parameter)
		req = req.WithContext(ctx)
	}

	handler.ServeHTTP(res, req)
}
