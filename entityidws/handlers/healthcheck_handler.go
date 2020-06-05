package handlers

import (
	"github.com/uvalib/entity-id-ws/entityidws/idservice"
	"net/http"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	// se have decided that because this is an external service, we do not report a healthcheck error
	// (which results in a service restart)
	_, message := idservice.GetStatus()
	encodeHealthCheckResponse(w, http.StatusOK, message)
}

//
// end of file
//
