package routego

import (
	"net/http"

	"github.com/hurtki/routego/internal/route"
)

// ===== http.Handler =====

func (r *Router) Get(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.GET)
}

func (r *Router) Post(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.POST)
}

func (r *Router) Put(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.PUT)
}

func (r *Router) Patch(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.PATCH)
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.DELETE)
}

func (r *Router) Head(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.HEAD)
}

func (r *Router) Options(path string, handler http.Handler) {
	r.routeSet.Add(path, handler.ServeHTTP, route.OPTIONS)
}

// ===== http.HandlerFunc =====

func (r *Router) GetFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.GET)
}

func (r *Router) PostFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.POST)
}

func (r *Router) PutFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.PUT)
}

func (r *Router) PatchFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.PATCH)
}

func (r *Router) DeleteFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.DELETE)
}

func (r *Router) HeadFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.HEAD)
}

func (r *Router) OptionsFunc(path string, handler http.HandlerFunc) {
	r.routeSet.Add(path, handler, route.OPTIONS)
}
