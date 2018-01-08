package handlers

import (
	"entityidws/idservice"
	"net/http"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	// update the statistics
	Statistics.RequestCount++
	Statistics.HeartbeatCount++

	status := idservice.GetStatus()
	message := ""
	encodeHealthCheckResponse(w, status, message)
}

//
// end of file
//
