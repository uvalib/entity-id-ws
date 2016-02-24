package main

import (
    "fmt"
    "time"
    "net/http"
    "strings"
    "bytes"
    "github.com/parnurzeal/gorequest"
)

// debug the http exchange
var debugHttp = false

//
// get entity details when provided a DOI
//
func GetDoi( doi string ) ( Entity, int ) {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.EzidServiceUrl, doi )
    fmt.Println( "GET URL:", url )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        Get( url  ).
        Timeout( time.Duration( config.EzidServiceTimeout ) * time.Second ).
        End( )
    fmt.Println( "Time:", time.Since( start ) )

    // check for errors
    if errs != nil {
        fmt.Println( "Errors:", errs )
        return blankEntity( ), http.StatusInternalServerError
    }

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
func CreateDoi( shoulder string, entity Entity ) ( Entity, int ) {

    // construct target URL
    url := fmt.Sprintf( "%s/shoulder/%s", config.EzidServiceUrl, shoulder )
    fmt.Println( "POST URL:", url )

    // construct the payload...
    body := makeBodyFromEntity( entity )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        SetBasicAuth( config.EzidUser, config.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.EzidServiceTimeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    fmt.Println( "Time:", time.Since( start ) )

    // check for errors
    if errs != nil {
        fmt.Println( "Errors:", errs )
        return blankEntity( ), http.StatusInternalServerError
    }

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
func UpdateDoi( entity Entity ) int {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.EzidServiceUrl, entity.Id )
    fmt.Println( "POST URL:", url )

    // construct the payload...
    body := makeBodyFromEntity( entity )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        SetBasicAuth( config.EzidUser, config.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.EzidServiceTimeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    fmt.Println( "Time:", time.Since( start ) )

    // check for errors
    if errs != nil {
        fmt.Println( "Errors:", errs )
        return http.StatusInternalServerError
    }

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
    url := fmt.Sprintf( "%s/id/%s", config.EzidServiceUrl, doi )
    fmt.Println( "DEL URL:", url )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        SetBasicAuth( config.EzidUser, config.EzidPassphrase ).
        Delete( url  ).
        Timeout( time.Duration( config.EzidServiceTimeout ) * time.Second ).
        End( )
    fmt.Println( "Time:", time.Since( start ) )

    // check for errors
    if errs != nil {
        fmt.Println( "Errors:", errs )
        return http.StatusInternalServerError
    }

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
    url := fmt.Sprintf( "%s/status", config.EzidServiceUrl )
    fmt.Println( "GET URL:", url )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        Get( url  ).
        Timeout( time.Duration( config.EzidServiceTimeout ) * time.Second ).
        End( )
    fmt.Println( "Time:", time.Since( start ) )

    // check for errors
    if errs != nil {
        fmt.Println( "Errors:", errs )
        return http.StatusInternalServerError
    }

    // check the body for errors
    if !statusIsOk( body ) {
        return http.StatusBadRequest
    }

    // all good...
    return http.StatusOK
}

//
// the response body consists of a set of CR separated lines containing tokens separated by
// a colon character
//
func makeEntityFromBody( body string ) Entity {

    //fmt.Println( "Response:", body )

    entity := blankEntity( )
    split := strings.Split( body, "\n" )
    for i := range split {
        tokens := strings.SplitN( split[ i ], ":", 2 )
        if len( tokens ) == 2 {
            s := strings.TrimSpace( tokens[ 1 ] )
            switch tokens[ 0 ] {
            case "success":
                entity.Id = s
            case "_target":
                entity.Url = s
            case "datacite.title":
                entity.Title = s
            case "datacite.publisher":
                entity.Publisher = s
            case "datacite.creator":
                entity.Creator = s
            case "datacite.publicationyear":
                entity.PubYear = s
            case "datacite.resourcetype":
                entity.ResourceType = s
            }
        }
    }
    return entity

}

func makeBodyFromEntity( entity Entity ) string {
    var buffer bytes.Buffer
    addBodyTerm( &buffer, "_target", entity.Url )
    addBodyTerm( &buffer, "datacite.title", entity.Title )
    addBodyTerm( &buffer, "datacite.publisher", entity.Publisher )
    addBodyTerm( &buffer, "datacite.creator", entity.Creator )
    addBodyTerm( &buffer, "datacite.publicationyear", entity.PubYear )
    addBodyTerm( &buffer, "datacite.resourcetype", entity.ResourceType )
    s := buffer.String( )
    //fmt.Println( "Payload:", s )
    return s
}

func addBodyTerm( buffer * bytes.Buffer, term string, value string ) {
    //fmt.Printf( "[%s] -> [%s]\n", term, value )

    if len( value ) != 0 {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, value ) )
    }
}

func blankEntity( ) Entity {
    return Entity{ }
}

func statusIsOk( body string ) bool {
    //fmt.Println( "Response:", body )
    return( strings.Index( body, "success:" ) == 0 )
}
