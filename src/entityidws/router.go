package main

import (
	"entityidws/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		handlers.HealthCheck,
	},

	Route{
		"VersionInfo",
		"GET",
		"/version",
		handlers.VersionInfo,
	},

	Route{
		"RuntimeInfo",
		"GET",
		"/runtime",
		handlers.RuntimeInfo,
	},

	Route{
		"StatsGet",
		"GET",
		"/statistics",
		handlers.StatsGet,
	},

	Route{
		"IdLookup",
		"GET",
		"/{doi:.*}",
		handlers.IdLookup,
	},

	Route{
		"IdCreate",
		"POST",
		"/{shoulder:.*}",
		handlers.IdCreate,
	},

	Route{
		"IdRevoke",
		"PUT",
		"/revoke/{doi:.*}",
		handlers.IdRevoke,
	},

	Route{
		"IdUpdate",
		"PUT",
		"/{doi:.*}",
		handlers.IdUpdate,
	},

	Route{
		"IdDelete",
		"DELETE",
		"/{doi:.*}",
		handlers.IdDelete,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
