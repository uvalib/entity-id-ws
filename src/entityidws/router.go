package main

import (
	"entityidws/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routeSlice []route

var routes = routeSlice{

	route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		handlers.HealthCheck,
	},

	route{
		"VersionInfo",
		"GET",
		"/version",
		handlers.VersionInfo,
	},

	route{
		"RuntimeInfo",
		"GET",
		"/runtime",
		handlers.RuntimeInfo,
	},

	route{
		"StatsGet",
		"GET",
		"/statistics",
		handlers.StatsGet,
	},

	route{
		"IDLookup",
		"GET",
		"/{doi:.*}",
		handlers.IDLookup,
	},

	route{
		"IDCreate",
		"POST",
		"/{shoulder:.*}",
		handlers.IDCreate,
	},

	route{
		"IDRevoke",
		"PUT",
		"/revoke/{doi:.*}",
		handlers.IDRevoke,
	},

	route{
		"IDUpdate",
		"PUT",
		"/{doi:.*}",
		handlers.IDUpdate,
	},

	route{
		"IDDelete",
		"DELETE",
		"/{doi:.*}",
		handlers.IDDelete,
	},
}

//
// NewRouter -- build and return the router
//
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

	// add the route for the expvars endpoint
	router.Handle("/debug/vars", http.DefaultServeMux )

	return router
}

//
// end of file
//
