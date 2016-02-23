package main

import (
   "encoding/json"
   "net/http"
   "github.com/gorilla/mux"
)

func TokenLookup( w http.ResponseWriter, r *http.Request ) {
    vars := mux.Vars( r )
    whom := vars[ "whom" ]
    what := vars[ "what" ]
    token := vars[ "token" ]

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    details, err := GetTokenDetails( token )

    if err != nil {
        w.WriteHeader( http.StatusInternalServerError )
        return
    }
   
    if ActivityIsOk( details, whom, what ) {
        status := http.StatusOK
        w.WriteHeader( status )
        if err := json.NewEncoder( w ).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
            panic( err )
        }
        return
    }

    // If this token is not OK then 403
    status := http.StatusForbidden
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
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