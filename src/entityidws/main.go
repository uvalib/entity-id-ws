package main

import (
   "entityidws/config"
   "entityidws/handlers"
   "entityidws/logger"
   "fmt"
   "log"
   "net/http"
)

func main() {

   logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

   // setup router and serve...
   router := NewRouter()
   log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Configuration.ServicePort), router))
}

//
// end of file
//