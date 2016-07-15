package ezid

import (
    "fmt"
    "log"
    "time"
    "net/http"
    "github.com/parnurzeal/gorequest"
    "entityidws/api"
    "entityidws/config"
)

const STATUS_RESERVED = "reserved"
const STATUS_PUBLIC = "public"

//
// get entity details when provided a DOI
//
func GetDoi( doi string ) ( api.Entity, int ) {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, doi )

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        Get( url  ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        log.Printf( "ERROR: service (%s) returns %s in %s\n", url, errs, duration )
        return blankEntity( ), http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    log.Printf( "Service (%s) returns http %d in %s\n", url, resp.StatusCode, duration )

    // check the body for errors
    if !statusIsOk( body ) {
        return blankEntity( ), http.StatusBadRequest
    }

    // all good...
    return makeEntityFromBody( body ), http.StatusOK
}

//
// Create a new entity; we may or may not have complete entity details
//
func CreateDoi( shoulder string, entity api.Entity ) ( api.Entity, int ) {

    // construct target URL
    url := fmt.Sprintf( "%s/shoulder/%s", config.Configuration.EzidServiceUrl, shoulder )

    // construct the payload, set the status to reserved
    body := makeDataciteBodyFromEntity( entity, STATUS_RESERVED )

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        log.Printf( "ERROR: service (%s) returns %s in %s\n", url, errs, duration )
        return blankEntity( ), http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    log.Printf( "Service (%s) returns http %d in %s\n", url, resp.StatusCode, duration )

    // check the body for errors
    if !statusIsOk( body ) {
        return blankEntity( ), http.StatusBadRequest
    }

    // all good...
    return makeEntityFromBody( body ), http.StatusOK
}

//
// Update an existing DOI to match the provided entity
//
func UpdateDoi( entity api.Entity ) int {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, entity.Id )

    // construct the payload...
    body := makeDataciteBodyFromEntity( entity, STATUS_PUBLIC )

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        log.Printf( "ERROR: service (%s) returns %s in %s\n", url, errs, duration )
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )
    log.Printf( "Service (%s) returns http %d in %s\n", url, resp.StatusCode, duration )

    // check the body for errors
    if !statusIsOk( body ) {
        return http.StatusBadRequest
    }

    // all good...
    return http.StatusOK
}

//
// delete entity details when provided a DOI
//
func DeleteDoi( doi string ) int {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, doi )

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Delete( url  ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        log.Printf( "ERROR: service (%s) returns %s in %s\n", url, errs, duration )
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )
    log.Printf( "Service (%s) returns http %d in %s\n", url, resp.StatusCode, duration )

    // check the body for errors
    if !statusIsOk( body ) {
        return http.StatusBadRequest
    }

    // all good...
    return http.StatusOK
}

//
// get the status of the endpoint
//
func GetStatus( ) int {

    // construct target URL
    url := fmt.Sprintf( "%s/status", config.Configuration.EzidServiceUrl )

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        Get( url  ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        log.Printf( "ERROR: service (%s) returns %s in %s\n", url, errs, duration )
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )
    log.Printf( "Service (%s) returns http %d in %s\n", url, resp.StatusCode, duration )

    // check the body for errors
    if !statusIsOk( body ) {
        return http.StatusBadRequest
    }

    // all good...
    return http.StatusOK
}