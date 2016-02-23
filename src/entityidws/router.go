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
        "IdLookup",
        "GET",
        "/entityid/{doi:.*}",
        IdLookup,
    },
    Route{
        "IdCreate",
        "POST",
        "/entityid/{shoulder:.*}",
        IdCreate,
    },
    Route{
        "IdUpdate",
        "PUT",
        "/entityid/{doi:.*}",
        IdUpdate,
    },
    Route{
        "IdDelete",
        "DEL",
        "/entityid/{doi:.*}",
        IdDelete,
    },

    Route{
        "HealthCheck",
        "GET",
        "/healthcheck",
        HealthCheck,
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
