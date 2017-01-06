package handlers

import (
    "net/http"
    "entityidws/ezid"
)

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    // update the statistics
    Statistics.RequestCount++
    Statistics.HeartbeatCount++

    status := ezid.GetStatus( )
    message := ""
    encodeHealthCheckResponse( w, status, message )
}