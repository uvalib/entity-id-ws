package main

import (
   "net/http"
   "github.com/gorilla/mux"
)

type Route struct {
   Name        string
   Method      string
   Pattern     string
   HandlerFunc http.HandlerFunc
}

type Routes [] Route

var routes = Routes{

    Route{
        "HealthCheck",
        "GET",
        "/healthcheck",
        HealthCheck,
    },

    Route{
        "GetVersion",
        "GET",
        "/version",
        GetVersion,
    },

    Route{
        "IdLookup",
        "GET",
        "/{doi:.*}",
        IdLookup,
    },

    Route{
        "IdCreate",
        "POST",
        "/{shoulder:.*}",
        IdCreate,
    },

    Route{
        "IdUpdate",
        "PUT",
        "/{doi:.*}",
        IdUpdate,
    },

    Route{
        "IdDelete",
        "DELETE",
        "/{doi:.*}",
        IdDelete,
    },

    Route{
        "Stats",
        "GET",
        "/statistics",
        Stats,
    },
}

func NewRouter( ) *mux.Router {

   router := mux.NewRouter().StrictSlash( true )
   for _, route := range routes {

      var handler http.Handler

      handler = route.HandlerFunc
      handler = Logger( handler, route.Name )

      router.
         Methods( route.Method ).
         Path( route.Pattern ).
         Name( route.Name ).
         Handler( handler )
   }

   return router
}
