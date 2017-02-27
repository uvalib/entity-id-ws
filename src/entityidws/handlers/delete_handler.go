package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "entityidws/ezid"
    "entityidws/authtoken"
    "entityidws/config"
)

func IdDelete( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]
    token := r.URL.Query( ).Get( "auth" )

    // update the statistics
    Statistics.RequestCount++
    Statistics.DeleteCount++

    // validate inbound parameters
    if parameterOK( doi ) == false || parameterOK( token ) == false {
        encodeStandardResponse( w, http.StatusBadRequest )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, "delete", token, config.Configuration.Timeout ) == false {
        encodeStandardResponse( w, http.StatusForbidden )
        return
    }

    status := ezid.DeleteDoi( doi )
    encodeStandardResponse( w, status )
}