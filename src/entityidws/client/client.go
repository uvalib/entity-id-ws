package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
    "entityidws/api"
    "encoding/json"
    "io"
    "io/ioutil"
)

var debugHttp = false
var serviceTimeout = 5

func HealthCheck( endpoint string ) int {

    url := fmt.Sprintf( "%s/healthcheck", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    return resp.StatusCode
}

func VersionCheck( endpoint string ) ( int, string ) {

    url := fmt.Sprintf( "%s/version", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
    SetDebug( debugHttp ).
    Get( url ).
    Timeout( time.Duration( serviceTimeout ) * time.Second ).
    End( )

    if errs != nil {
        return http.StatusInternalServerError, ""
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.VersionResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, ""
    }

    return resp.StatusCode, r.Version
}

func RuntimeCheck( endpoint string ) ( int, * api.RuntimeResponse ) {

    url := fmt.Sprintf( "%s/runtime", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
            SetDebug( false ).
            Get( url  ).
            Timeout( time.Duration( serviceTimeout ) * time.Second ).
            End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.RuntimeResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, &r
}

func Statistics( endpoint string ) ( int, * api.Statistics ) {

    url := fmt.Sprintf( "%s/statistics", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.StatisticsResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, &r.Details
}

func Get( endpoint string, doi string, token string ) ( int, * api.Entity ) {

    url := fmt.Sprintf( "%s/%s?auth=%s", endpoint, doi, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError, nil
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}

func Create( endpoint string, shoulder string, token string ) ( int, * api.Entity ) {

    url := fmt.Sprintf("%s/%s?auth=%s", endpoint, shoulder, token)
    //fmt.Printf( "%s\n", url )

    entity := api.Entity{ Title : "my title", Url: "http://google.com" }

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Post( url ).
       Send( entity ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       Set( "Content-Type", "application/json" ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}

func Update( endpoint string, entity api.Entity, token string ) int {

    url := fmt.Sprintf("%s/%s?auth=%s", endpoint, entity.Id, token )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Put( url ).
       Send( entity ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       Set( "Content-Type", "application/json" ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    return resp.StatusCode
}

func Delete( endpoint string, doi string, token string ) int {

    url := fmt.Sprintf("%s/%s?auth=%s", endpoint, doi, token )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Delete( url ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    return resp.StatusCode
}

//
// revoke an entity when provided a DOI
//
func Revoke( endpoint string, doi string, token string ) int {

    // construct target URL
    url := fmt.Sprintf("%s/revoke/%s?auth=%s", endpoint, doi, token )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Put( url ).
       Timeout( time.Duration( serviceTimeout ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    return resp.StatusCode
}
