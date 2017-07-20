package handlers

import (
	"entityidws/api"
	"net/http"
)

var Statistics = api.Statistics{}

func StatsGet(w http.ResponseWriter, r *http.Request) {
	// create response
	encodeStatsResponse(w, Statistics)
}
