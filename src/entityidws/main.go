package main

import (
    "fmt"
    "log"
    "net/http"
    "entityidws/api"
    "entityidws/config"
)

var statistics = api.Statistics{ }

func main( ) {

    log.Printf( "===> %s version: '%s' <===", config.Configuration.ServiceName, Version( ) )

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.ServicePort ), router ) )
}

