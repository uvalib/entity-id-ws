package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "entityidws/api"
    "entityidws/ezid"
    "log"
)

func IdLookup( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]

    // update the statistics
    statistics.RequestCount++
    statistics.LookupCount++

    entity, status := ezid.GetDoi( doi )
    respondWithDetails( w, status, entity )
}

func IdCreate( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    shoulder := vars[ "shoulder" ]

    decoder := json.NewDecoder( r.Body )
    entity := api.Entity{ }

    // update the statistics
    statistics.RequestCount++
    statistics.CreateCount++

    if err := decoder.Decode( &entity ); err != nil {
        respond( w, http.StatusBadRequest )
        return
    }

    entity, status := ezid.CreateDoi( shoulder, entity )
    respondWithDetails( w, status, entity )
}

func IdUpdate( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]

    decoder := json.NewDecoder( r.Body )
    entity := api.Entity{ }

    // update the statistics
    statistics.RequestCount++
    statistics.UpdateCount++

    if err := decoder.Decode( &entity ); err != nil {
        respond( w, http.StatusBadRequest )
        return
    }

    entity.Id = doi
    status := ezid.UpdateDoi( entity )
    respond( w, status )
}

func IdDelete( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]

    // update the statistics
    statistics.RequestCount++
    statistics.DeleteCount++

    status := ezid.DeleteDoi( doi )
    respond( w, status )
}

func Stats( w http.ResponseWriter, r *http.Request ) {

    status := http.StatusOK

    jsonResponse( w )
    w.WriteHeader( status )

    if err := json.NewEncoder(w).Encode( api.StatisticsResponse { Status: status, Message: http.StatusText( status ), Details: statistics } ); err != nil {
        log.Fatal( err )
    }
}

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    // update the statistics
    statistics.RequestCount++
    statistics.HeartbeatCount++

    status := ezid.GetStatus( )
    healthy := status == http.StatusOK
    message := ""

    jsonResponse( w )
    w.WriteHeader( status )

    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func respond( w http.ResponseWriter, status int ) {

    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        log.Fatal( err )
    }
}

func respondWithDetails( w http.ResponseWriter, status int, entity api.Entity ) {

    jsonResponse( w )
    w.WriteHeader( status )
    if status == http.StatusOK {
        if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: http.StatusText( status ), Details: entity } ); err != nil {
            log.Fatal( err )
        }
    } else {
        if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: http.StatusText( status ) } ); err != nil {
            log.Fatal( err )
        }
    }
}

func jsonResponse( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}