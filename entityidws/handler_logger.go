package main

import (
	"github.com/uvalib/entity-id-ws/entityidws/logger"
	"fmt"
	"net/http"
	"time"
)

//
// HandlerLogger -- middleware handler
//
func HandlerLogger(inner http.Handler, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Log(fmt.Sprintf(
			"%s (%s) -> method %s, time %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))
	})
}

//
// end of file
//
