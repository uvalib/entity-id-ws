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
        "/entityid/{namespace}/{id}",
        IdLookup,
    },
    Route{
        "IdCreate",
        "POST",
        "/entityid",
        IdCreate,
    },
    Route{
        "IdUpdate",
        "PUT",
        "/entityid",
        IdUpdate,
    },
    Route{
        "IdDelete",
        "DEL",
        "/entityid/{namespace}/{id}",
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
