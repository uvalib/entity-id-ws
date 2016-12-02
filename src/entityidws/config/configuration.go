package config

import (
    "flag"
    "fmt"
    "entityidws/logger"
    "strings"
)

type Config struct {
    ServicePort        string
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
    flag.StringVar( &c.EzidServiceUrl, "ezidurl", "https://ezid.cdlib.org", "The EZID service URL" )
    flag.IntVar( &c.EzidServiceTimeout, "timeout", 10, "The service timeout (in seconds)")
    flag.StringVar( &c.EzidUser, "eziduser", "apitest", "The EZID service username" )
    flag.StringVar( &c.EzidPassphrase, "ezidpassword", "apitest", "The EZID service passphrase")
    flag.StringVar( &c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
    flag.BoolVar( &c.Debug, "debug", false, "Enable debugging")

    flag.Parse( )

    logger.Log( fmt.Sprintf( "ServicePort:        %s", c.ServicePort ) )
    logger.Log( fmt.Sprintf( "EzidServiceUrl:     %s", c.EzidServiceUrl ) )
    logger.Log( fmt.Sprintf( "EzidServiceTimeout: %d", c.EzidServiceTimeout ) )
    logger.Log( fmt.Sprintf( "EzidUser:           %s", c.EzidUser ) )
    logger.Log( fmt.Sprintf( "EzidPassphrase:     %s", strings.Repeat( "*", len( c.EzidPassphrase ) ) ) )
    logger.Log( fmt.Sprintf( "AuthTokenEndpoint   %s", c.AuthTokenEndpoint ) )
    logger.Log( fmt.Sprintf( "Debug               %t", c.Debug ) )

    return c
}
