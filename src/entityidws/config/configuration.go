package config

import (
    "flag"
    "log"
)
type Config struct {
    ServicePort        string
    AuthorizerUrl      string
    EzidServiceUrl     string
    EzidServiceTimeout int
    EzidUser           string
    EzidPassphrase     string
    AuthTokenEndpoint  string
    Debug              bool
}

var Configuration = LoadConfig( )

func LoadConfig( ) Config {

    c := Config{ }

    // process command line flags and setup configuration
    flag.StringVar( &c.ServicePort, "port", "8080", "The service listen port" )
    flag.StringVar( &c.AuthorizerUrl, "authurl", "docker1.lib.virginia.edu:8200", "The authorizer service hostname:port" )
    flag.StringVar( &c.EzidServiceUrl, "ezidurl", "https://ezid.cdlib.org", "The EZID service URL" )
    flag.IntVar( &c.EzidServiceTimeout, "timeout", 10, "The service timeout (in seconds)")
    flag.StringVar( &c.EzidUser, "eziduser", "apitest", "The EZID service username" )
    flag.StringVar( &c.EzidPassphrase, "ezidpassword", "apitest", "The EZID service passphrase")
    flag.StringVar( &c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
    flag.BoolVar( &c.Debug, "debug", false, "Enable debugging")

    flag.Parse( )

    log.Printf( "ServicePort:        %s", c.ServicePort )
    log.Printf( "AuthorizerUrl:      %s", c.AuthorizerUrl )
    log.Printf( "EzidServiceUrl:     %s", c.EzidServiceUrl )
    log.Printf( "EzidServiceTimeout: %d", c.EzidServiceTimeout )
    log.Printf( "EzidUser:           %s", c.EzidUser )
    log.Printf( "EzidPassphrase:     %s", c.EzidPassphrase )
    log.Printf( "AuthTokenEndpoint   %s", c.AuthTokenEndpoint )
    log.Printf( "Debug               %t", c.Debug )

    return c
}
