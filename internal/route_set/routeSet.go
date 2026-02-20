package route_set

import (
	"net/http"

	"github.com/hurtki/routego/internal/route"
)

// RouteSet is for storing path to handler map
type RouteSet struct {
	routes []route.Route
}

func NewRouteSet() RouteSet {
	return RouteSet{routes: []route.Route{}}
}

// Add(pattern, handler) creates a new route in routeSet with a pattern
// pattern can contain one of parameters: {num}, {string}
// pattern can contain strict parts: /tasks/, /api/users/
// examples: 'api/users/3', 'post/{string}'
func (s *RouteSet) Add(path string, handler http.HandlerFunc, httpMethod route.HttpMethod) {
	route := route.NewRoute(path, handler, httpMethod)
	s.routes = append(s.routes, route)
}

// Handler(path string) tries to match given path specified with Add() routes
// returns a handler, url parameter(if there is no then nil), bool (true, if found)
func (s *RouteSet) Handler(path string, httpMethod route.HttpMethod) (http.Handler, any, bool) {
	for _, route := range s.routes {

		matches, parameter := route.Match(path, httpMethod)
		if matches {
			return route.Handler, parameter, true
		}
	}

	return nil, nil, false
}
