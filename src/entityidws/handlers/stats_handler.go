package handlers

import (
    "net/http"
    "entityidws/api"
)

var Statistics = api.Statistics{ }

func StatsGet( w http.ResponseWriter, r *http.Request ) {
    // create response
    encodeStatsResponse( w, Statistics )
}