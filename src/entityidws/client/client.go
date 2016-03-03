package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
    "entityidws/api"
    "encoding/json"
)

func HealthCheck( endpoint string ) int {

    url := fmt.Sprintf( "%s/healthcheck", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}

func Get( endpoint string, doi string, token string ) ( int, * api.Entity ) {

    url := fmt.Sprintf( "%s/entityid/%s?auth=%s", endpoint, doi, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}

func Create( endpoint string, shoulder string, token string ) ( int, * api.Entity ) {

    url := fmt.Sprintf("%s/entityid/%s?auth=%s", endpoint, shoulder, token)
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( false ).
       Post( url ).
       Send( api.Entity{ Title : "my title" } ).
       Timeout( time.Duration( 5 ) * time.Second ).
       Set( "Content-Type", "application/json" ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}

func Update( endpoint string, entity api.Entity, token string ) int {

    url := fmt.Sprintf("%s/entityid/%s?auth=%s", endpoint, entity.Id, token )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Put( url ).
       Send( entity ).
       Timeout( time.Duration( 5 ) * time.Second ).
       Set( "Content-Type", "application/json" ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}

