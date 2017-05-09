package ezid

import (
    "fmt"
    "time"
    "net/http"
    "github.com/parnurzeal/gorequest"
    "entityidws/api"
    "entityidws/config"
    "entityidws/logger"
    "io"
    "io/ioutil"
)

// status for the EZID objects
const STATUS_PUBLIC = "public"
const STATUS_RESERVED = "reserved"
const STATUS_UNAVAILABLE = "unavailable|withdrawn by Library"

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
        Timeout( time.Duration( config.Configuration.Timeout ) * time.Second ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: service (%s) returns %s in %s", url, errs, duration ) )
        return blankEntity( ), http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Service (%s) returns http %d in %s", url, resp.StatusCode, duration ) )

    // check the body for errors
    if !statusIsOk( body ) {
        logger.Log( fmt.Sprintf( "Error response body: [%s]", body ) )
        return blankEntity( ), http.StatusBadRequest
    }

    // all good...
    return makeEntityFromBody( body ), http.StatusOK
}

//
// Create a new entity; we may or may not have complete entity details
//
func CreateDoi( shoulder string, entity api.Entity, status string ) ( api.Entity, int ) {

    // log if necessary
    logEntity( entity )

    // construct target URL
    url := fmt.Sprintf( "%s/shoulder/%s", config.Configuration.EzidServiceUrl, shoulder )

    // build the request body
    body, err := makeBodyFromEntity( entity, status )

    // check for errors
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: creating service payload %s", err ) )
        return blankEntity( ), http.StatusBadRequest
    }

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.Configuration.Timeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: service (%s) returns %s in %s", url, errs, duration ) )
        return blankEntity( ), http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Service (%s) returns http %d in %s", url, resp.StatusCode, duration ) )

    // check the body for errors
    if !statusIsOk( body ) {
        logger.Log( fmt.Sprintf( "Error response body: [%s]", body ) )
        return blankEntity( ), http.StatusBadRequest
    }

    // all good...
    return makeEntityFromBody( body ), http.StatusOK
}

//
// Update an existing DOI to match the provided entity
//
func UpdateDoi( entity api.Entity, status string ) int {

    // log if necessary
    logEntity( entity )

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, entity.Id )

    // build the request body
    body, err := makeBodyFromEntity( entity, status )

    // check for errors
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: creating service payload %s", err ) )
        return http.StatusBadRequest
    }

    // issue the request
    start := time.Now( )
    resp, body, errs := gorequest.New( ).
        SetDebug( config.Configuration.Debug ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.Configuration.Timeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: service (%s) returns %s in %s", url, errs, duration ) )
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Service (%s) returns http %d in %s", url, resp.StatusCode, duration ) )

    // check the body for errors
    if !statusIsOk( body ) {
        logger.Log( fmt.Sprintf( "Error response body: [%s]", body ) )
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
        Timeout( time.Duration( config.Configuration.Timeout ) * time.Second ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: service (%s) returns %s in %s", url, errs, duration ) )
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Service (%s) returns http %d in %s", url, resp.StatusCode, duration ) )

    // check the body for errors
    if !statusIsOk( body ) {
        logger.Log( fmt.Sprintf( "Error response body: [%s]", body ) )
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
        Timeout( time.Duration( config.Configuration.Timeout ) * time.Second ).
        End( )
    duration := time.Since( start )

    // check for errors
    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: service (%s) returns %s in %s", url, errs, duration ) )
        return http.StatusInternalServerError
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Service (%s) returns http %d in %s", url, resp.StatusCode, duration ) )

    // check the body for errors
    if !statusIsOk( body ) {
        logger.Log( fmt.Sprintf( "Error response body: [%s]", body ) )
        return http.StatusBadRequest
    }

    // all good...
    return http.StatusOK
}