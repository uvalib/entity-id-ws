package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
//    "fmt"
)

func IdLookup( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]

    // update the statistics
    statistics.RequestCount++
    statistics.LookupCount++

    entity, status := GetDoi( doi )
    respondWithDetails( w, status, entity )
}

func IdCreate( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    shoulder := vars[ "shoulder" ]

    decoder := json.NewDecoder( r.Body )
    entity := Entity{ }

    // update the statistics
    statistics.RequestCount++
    statistics.CreateCount++

    if err := decoder.Decode( &entity ); err != nil {
        respond( w, http.StatusBadRequest )
        return
    }

    entity, status := CreateDoi( shoulder, entity )
    respondWithDetails( w, status, entity )
}

func IdUpdate( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]

    decoder := json.NewDecoder( r.Body )
    entity := Entity{ }

    // update the statistics
    statistics.RequestCount++
    statistics.UpdateCount++

    if err := decoder.Decode( &entity ); err != nil {
        respond( w, http.StatusBadRequest )
        return
    }

    entity.Id = doi
    status := UpdateDoi( entity )
    respond( w, status )
}

func IdDelete( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    doi := vars[ "doi" ]

    // update the statistics
    statistics.RequestCount++
    statistics.DeleteCount++

    status := DeleteDoi( doi )
    respond( w, status )
}

func Stats( w http.ResponseWriter, r *http.Request ) {

    status := http.StatusOK

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( status )

    if err := json.NewEncoder(w).Encode( StatisticsResponse { Status: status, Message: http.StatusText( status ), Details: statistics } ); err != nil {
        panic(err)
    }
}

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    status := GetStatus( )
    healthy := status == http.StatusOK
    message := ""

    // update the statistics
    statistics.RequestCount++
    statistics.HeartbeatCount++

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( status )

    if err := json.NewEncoder(w).Encode( HealthCheckResponse { CheckType: HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        panic(err)
    }
}

func respond( w http.ResponseWriter, status int ) {

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        panic(err)
    }
}

func respondWithDetails( w http.ResponseWriter, status int, entity Entity ) {

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( status )
    if status == http.StatusOK {
        if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ), Details: entity } ); err != nil {
            panic(err)
        }
    } else {
        if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
            panic(err)
        }
    }
}