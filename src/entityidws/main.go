package main

import (
   "fmt"
   "log"
   "net/http"
   "flag"
)

var config = Configuration{ }
var statistics = Statistics{ }

func main( ) {

	// process command line flags and setup configuration
	flag.StringVar( &config.ServicePort, "port", "8080", "The service listen port" )
	flag.StringVar( &config.AuthorizerUrl, "authurl", "docker1.lib.virginia.edu:8200", "The authorizer service hostname:port" )
    flag.StringVar( &config.EzidServiceUrl, "ezidurl", "https://ezid.cdlib.org", "The EZID service URL" )
    flag.IntVar( &config.EzidServiceTimeout, "timeout", 10, "The service timeout (in seconds)")
	flag.StringVar( &config.EzidUser, "eziduser", "apitest", "The EZID service username" )
	flag.StringVar( &config.EzidPassphrase, "ezidpassword", "apitest", "The EZID service passphrase")

	flag.Parse()

	log.Printf( "ServicePort:        %s", config.ServicePort )
	log.Printf( "AuthorizerUrl:      %s", config.AuthorizerUrl )
    log.Printf( "EzidServiceUrl:     %s", config.EzidServiceUrl )
    log.Printf( "EzidServiceTimeout: %d", config.EzidServiceTimeout )
	log.Printf( "EzidUser:           %s", config.EzidUser )
	log.Printf( "EzidPassphrase:     %s", config.EzidPassphrase )

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.ServicePort ), router ) )
}

