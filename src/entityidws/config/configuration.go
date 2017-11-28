package config

import (
	"entityidws/logger"
	"flag"
	"fmt"
	"strings"
)

//
// Config -- our configuration structure
//
type Config struct {
	ServicePort       string
	EzidServiceURL    string
	EzidUser          string
	EzidPassphrase    string
	AuthTokenEndpoint string
	Timeout           int
	Debug             bool
}

//
// Configuration -- our configuration instance
//
var Configuration = loadConfig()

func loadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.EzidServiceURL, "ezidurl", "https://ezid.cdlib.org", "The EZID service URL")
	flag.StringVar(&c.EzidUser, "eziduser", "apitest", "The EZID service username")
	flag.StringVar(&c.EzidPassphrase, "ezidpassword", "apitest", "The EZID service passphrase")
	flag.StringVar(&c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
	flag.IntVar(&c.Timeout, "timeout", 15, "The external service timeout in seconds")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("ServicePort:        %s", c.ServicePort))
	logger.Log(fmt.Sprintf("EzidServiceURL:     %s", c.EzidServiceURL))
	logger.Log(fmt.Sprintf("EzidUser:           %s", c.EzidUser))
	logger.Log(fmt.Sprintf("EzidPassphrase:     %s", strings.Repeat("*", len(c.EzidPassphrase))))
	logger.Log(fmt.Sprintf("AuthTokenEndpoint   %s", c.AuthTokenEndpoint))
	logger.Log(fmt.Sprintf("Timeout:            %d", c.Timeout))
	logger.Log(fmt.Sprintf("Debug               %t", c.Debug))

	return c
}

//
// end of file
//
