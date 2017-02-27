package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "entityidws/ezid"
    "entityidws/authtoken"
    "entityidws/config"
)

func IdLookup( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]
    token := r.URL.Query( ).Get( "auth" )

    // update the statistics
    Statistics.RequestCount++
    Statistics.LookupCount++

    // validate inbound parameters
    if parameterOK( doi ) == false || parameterOK( token ) == false {
        encodeStandardResponse( w, http.StatusBadRequest )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, "lookup", token, config.Configuration.Timeout ) == false {
        encodeStandardResponse( w, http.StatusForbidden )
        return
    }

    entity, status := ezid.GetDoi( doi )
    encodeDetailsResponse( w, status, entity )
}