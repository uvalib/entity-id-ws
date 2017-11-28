package handlers

import (
	"entityidws/ezid"
	"net/http"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	// update the statistics
	Statistics.RequestCount++
	Statistics.HeartbeatCount++

	status := ezid.GetStatus()
	message := ""
	encodeHealthCheckResponse(w, status, message)
}

//
// end of file
//
