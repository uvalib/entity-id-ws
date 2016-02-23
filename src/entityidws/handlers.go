package main

import (
   "encoding/json"
   "net/http"
   "github.com/gorilla/mux"
    "fmt"
)

func IdLookup( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    namespace := vars[ "namespace" ]
    id := vars[ "id" ]

    doi := fmt.Sprintf( "%s/%s", namespace, id )

    // update the statistics
    statistics.RequestCount++
    statistics.LookupCount++

    entity, status := GetByDoi( doi )
    respond( w, status, entity )
}

func IdCreate( w http.ResponseWriter, r *http.Request ) {
//    vars := mux.Vars( r )
//    id := vars[ "id" ]

    // update the statistics
    statistics.RequestCount++
    statistics.CreateCount++

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )

    // If this token is not OK then 403
    status := http.StatusForbidden
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        panic(err)
    }
}

func IdUpdate( w http.ResponseWriter, r *http.Request ) {
//    vars := mux.Vars( r )
//    id := vars[ "id" ]

    // update the statistics
    statistics.RequestCount++
    statistics.UpdateCount++

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )

    // If this token is not OK then 403
    status := http.StatusForbidden
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        panic(err)
    }
}

func IdDelete( w http.ResponseWriter, r *http.Request ) {
//    vars := mux.Vars( r )
//    id := vars[ "id" ]

    // update the statistics
    statistics.RequestCount++
    statistics.DeleteCount++

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )

    // If this token is not OK then 403
    status := http.StatusForbidden
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        panic(err)
    }
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

    healthy := true
    message := ""

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( http.StatusOK )

    if err := json.NewEncoder(w).Encode( HealthCheckResponse { CheckType: HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        panic(err)
    }
}

func respond( w http.ResponseWriter, status int, entity Entity ) {
    
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