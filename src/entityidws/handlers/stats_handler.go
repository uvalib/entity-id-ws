package handlers

import (
   "entityidws/api"
   "net/http"
)

//
// Statistics -- the statistics instance
//
var Statistics = api.Statistics{}

//
// StatsGet -- get the current statistics
//
func StatsGet(w http.ResponseWriter, r *http.Request) {
   // create response
   encodeStatsResponse(w, Statistics)
}

//
// end of file
//
