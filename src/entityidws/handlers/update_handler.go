package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "entityidws/api"
    "entityidws/ezid"
    "entityidws/authtoken"
    "entityidws/config"
)

func IdUpdate( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]
    token := r.URL.Query( ).Get( "auth" )

    // update the statistics
    Statistics.RequestCount++
    Statistics.UpdateCount++

    // validate inbound parameters
    if parameterOK( doi ) == false || parameterOK( token ) == false {
        encodeStandardResponse( w, http.StatusBadRequest )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, "update", token, config.Configuration.Timeout ) == false {
        encodeStandardResponse( w, http.StatusForbidden )
        return
    }

    decoder := json.NewDecoder( r.Body )
    request := api.Request{ }

    if err := decoder.Decode( &request ); err != nil {
        encodeStandardResponse( w, http.StatusBadRequest )
        return
    }

    request.Id = doi
    status := ezid.UpdateDoi( request, ezid.STATUS_PUBLIC )
    encodeStandardResponse( w, status )
}