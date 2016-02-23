package main

import (
   "fmt"
   "log"
   "net/http"
   "flag"
)

var config = Configuration{ }

func main( ) {

	// process command line flags and serup configuration
	flag.StringVar( &config.ServicePort, "port", "8080", "The service listen port" )
	flag.StringVar( &config.DbHost, "dbhost", "mysqldev.lib.virginia.edu", "The database server hostname:port" )
    flag.StringVar( &config.DbName, "dbname", "authtoken_development", "The database name" )
	flag.StringVar( &config.DbUser, "dbuser", "dbuser", "The database username" )
	flag.StringVar( &config.DbPassphrase, "dbpassword", "dbpassword", "The database passphrase")

	flag.Parse()

	log.Printf( "ServicePort:  %s", config.ServicePort )
	log.Printf( "DbHost:       %s", config.DbHost )
    log.Printf( "DbName:       %s", config.DbName )
	log.Printf( "DbUser:       %s", config.DbUser )
	log.Printf( "DbPassphrase: %s", config.DbPassphrase )

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.ServicePort ), router ) )
}

