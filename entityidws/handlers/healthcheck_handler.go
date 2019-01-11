package handlers

import (
	"github.com/uvalib/entity-id-ws/entityidws/idservice"
	"net/http"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	status := idservice.GetStatus()
	message := ""
	encodeHealthCheckResponse(w, status, message)
}

//
// end of file
//