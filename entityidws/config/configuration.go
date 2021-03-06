package config

import (
	"flag"
	"fmt"
	"github.com/uvalib/entity-id-ws/entityidws/logger"
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
	SharedSecret        string
	ServiceTimeout      int
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
	flag.StringVar(&c.IDServiceURL, "idserviceurl", "https://not.configured.org", "The ID service URL")
	flag.StringVar(&c.IDServiceUser, "idserviceuser", "default", "The ID service username")
	flag.StringVar(&c.IDServicePassphrase, "idservicepasswd", "default", "The ID service passphrase")
	flag.StringVar(&c.SharedSecret, "secret", "", "The JWT shared secret")
	flag.IntVar(&c.ServiceTimeout, "timeout", 15, "The external service timeout in seconds")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("INFO: ServicePort:         %s", c.ServicePort))
	logger.Log(fmt.Sprintf("INFO: IDServiceURL:        %s", c.IDServiceURL))
	logger.Log(fmt.Sprintf("INFO: IDServiceUser:       %s", c.IDServiceUser))
	logger.Log(fmt.Sprintf("INFO: IDServicePassphrase: %s", strings.Repeat("*", len(c.IDServicePassphrase))))
	logger.Log(fmt.Sprintf("INFO: SharedSecret:        %s", strings.Repeat("*", len(c.SharedSecret))))
	logger.Log(fmt.Sprintf("INFO: ServiceTimeout:      %d", c.ServiceTimeout))
	logger.Log(fmt.Sprintf("INFO: Debug:               %t", c.Debug))

	return c
}

//
// end of file
//
