package route

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// route stores sample of url route
// and matches it with a handler
type Route struct {
	parts   []routePart
	method  HttpMethod
	Handler http.HandlerFunc
}

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PATCH   HttpMethod = "PATCH"
	PUT     HttpMethod = "PUT"
	DELETE  HttpMethod = "DELETE"
	OPTIONS HttpMethod = "OPTIONS"
	HEAD    HttpMethod = "HEAD"
)

// NewRoute(pattern, handler) creates a new route with a pattern
// pattern can contain one of parameters: {num}, {string}
// pattern can contain strict parts: /tasks/, /api/users/
// examples: 'api/users/3', 'post/{string}'
func NewRoute(pattern string, handler http.HandlerFunc, method HttpMethod) Route {
	patternParts := strings.FieldsFunc(pattern, func(r rune) bool {
		return r == '/'
	})

	routeParts := []routePart{}
	wasParameter := false
	for _, patternPart := range patternParts {
		if patternPart == "" {
			continue
		}

		routePart, err := NewRoutePart(patternPart)
		if err != nil {
			panic(err)
		}
		if wasParameter && !routePart.Strict {
			panic(ErrTwoParameteresInOneRoute)
		}
		if !routePart.Strict {
			wasParameter = true
		}
		routeParts = append(routeParts, routePart)
	}
	return Route{
		parts:   routeParts,
		Handler: handler,
		method:  method,
	}
}

// Match(path) checks if given path matches pattern of the route
func (r *Route) Match(path string, httpMethod HttpMethod) (bool, any) {
	if r.method != httpMethod {
		fmt.Printf("bad httpMethod, expected: %s, real: %s\n", r.method, httpMethod)
		return false, nil
	}
	pathParts := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	var resParameter any = nil

	if len(pathParts) != len(r.parts) {
		fmt.Println("wrong length")
		return false, nil
	}

	for i, routePart := range r.parts {
		pathPart := pathParts[i]
		if pathPart == "" {
			continue
		}

		ok, parameter := routePart.Compare(pathPart)
		if !ok {
			fmt.Printf("bad routepart, index %s\n", strconv.Itoa(i))
			return false, nil
		}

		if parameter != nil {
			resParameter = parameter
		}
	}

	return true, resParameter
}
