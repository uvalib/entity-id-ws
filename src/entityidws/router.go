package main

import (
	"entityidws/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
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

	// add the route for the prometheus metrics
	router.Handle("/metrics", HandlerLogger( promhttp.Handler( ), "promhttp.Handler" ) )

	// then add the remaining routes
	for _, route := range routes {

		var handler http.Handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)
		handler = prometheus.InstrumentHandler( route.Name, handler )

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

//
// end of file
//
