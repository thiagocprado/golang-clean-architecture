package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiagocprado/golang-api-structure/pkg/middlewares"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}

type Router struct {
	Router *mux.Router
}

func New() *Router {
	return &Router{Router: mux.NewRouter()}
}

func (r *Router) InjectRoutes(rt []Route) {
	for _, route := range rt {
		if route.RequestAuthentication {
			r.Router.HandleFunc(
				route.URI,
				middlewares.Authorization(route.Function),
			).Methods(route.Method)
			continue
		}

		r.Router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
}
