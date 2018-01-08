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
	ServicePort         string
	IDServiceURL        string
	IDServiceUser       string
	IDServicePassphrase string
	AuthTokenEndpoint   string
	Timeout             int
	Debug               bool
}

//
// Configuration -- our configuration instance
//
var Configuration = loadConfig()

func loadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.IDServiceURL, "idserviceurl", "https://ezid.cdlib.org", "The ID service URL")
	flag.StringVar(&c.IDServiceUser, "idserviceuser", "default", "The ID service username")
	flag.StringVar(&c.IDServicePassphrase, "idservicepasswd", "default", "The ID service passphrase")
	flag.StringVar(&c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
	flag.IntVar(&c.Timeout, "timeout", 15, "The external service timeout in seconds")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("ServicePort:         %s", c.ServicePort))
	logger.Log(fmt.Sprintf("IDServiceURL:        %s", c.IDServiceURL))
	logger.Log(fmt.Sprintf("IDServiceUser:       %s", c.IDServiceUser))
	logger.Log(fmt.Sprintf("IDServicePassphrase: %s", strings.Repeat("*", len(c.IDServicePassphrase))))
	logger.Log(fmt.Sprintf("AuthTokenEndpoint    %s", c.AuthTokenEndpoint))
	logger.Log(fmt.Sprintf("Timeout:             %d", c.Timeout))
	logger.Log(fmt.Sprintf("Debug                %t", c.Debug))

	return c
}

//
// end of file
//
