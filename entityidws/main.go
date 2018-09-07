package main

import (
	"github.com/uvalib/entity-id-ws/entityidws/config"
	"github.com/uvalib/entity-id-ws/entityidws/handlers"
	"github.com/uvalib/entity-id-ws/entityidws/logger"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// setup router and server...
	serviceTimeout := 15 * time.Second
	router := NewRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Configuration.ServicePort),
		Handler:      router,
		ReadTimeout:  serviceTimeout,
		WriteTimeout: serviceTimeout,
	}
	log.Fatal(server.ListenAndServe())
}

//
// end of file
//
