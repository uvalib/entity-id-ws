package ezid

import (
    "fmt"
    "time"
    "net/http"
    "strings"
    "bytes"
    "github.com/parnurzeal/gorequest"
    "entityidws/api"
    "entityidws/config"
)

// debug the http exchange
var debugHttp = false

//
// get entity details when provided a DOI
//
func GetDoi( doi string ) ( api.Entity, int ) {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, doi )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        Get( url  ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        End( )
    fmt.Printf( "GET: %s (%s)\n", url, time.Since( start ) )

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
func CreateDoi( shoulder string, entity api.Entity ) ( api.Entity, int ) {

    // construct target URL
    url := fmt.Sprintf( "%s/shoulder/%s", config.Configuration.EzidServiceUrl, shoulder )

    // construct the payload...
    body := makeBodyFromEntity( entity )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    fmt.Printf( "POST: %s (%s)\n", url, time.Since( start ) )

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
func UpdateDoi( entity api.Entity ) int {

    // construct target URL
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, entity.Id )

    // construct the payload...
    body := makeBodyFromEntity( entity )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Post( url  ).
        Send( body ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        Set( "Content-Type", "text/plain" ).
        End( )
    fmt.Printf( "POST %s (%s)\n", url, time.Since( start ) )

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
    url := fmt.Sprintf( "%s/id/%s", config.Configuration.EzidServiceUrl, doi )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        SetBasicAuth( config.Configuration.EzidUser, config.Configuration.EzidPassphrase ).
        Delete( url  ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        End( )
    fmt.Printf( "DELETE: %s (%s)\n", url, time.Since( start ) )

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
    url := fmt.Sprintf( "%s/status", config.Configuration.EzidServiceUrl )

    // issue the request
    start := time.Now( )
    _, body, errs := gorequest.New( ).
        SetDebug( debugHttp ).
        Get( url  ).
        Timeout( time.Duration( config.Configuration.EzidServiceTimeout ) * time.Second ).
        End( )
    fmt.Printf( "GET: %s (%s)\n", url, time.Since( start ) )

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
func makeEntityFromBody( body string ) api.Entity {

    //fmt.Println( "Response:", body )

    entity := blankEntity( )
    split := strings.Split( body, "\n" )
    for i := range split {
        tokens := strings.SplitN( split[ i ], ":", 2 )
        if len( tokens ) == 2 {
            s := strings.TrimSpace( tokens[ 1 ] )
            switch tokens[ 0 ] {
            case "success":
                entity.Id = strings.TrimSpace( strings.Split( s, "|" )[ 0 ] )
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

func makeBodyFromEntity( entity api.Entity ) string {
    var buffer bytes.Buffer
    addBodyTerm( &buffer, "_target", entity.Url, "https://virginia.edu" )
    addBodyTerm( &buffer, "datacite.title", entity.Title, "empty" )
    addBodyTerm( &buffer, "datacite.publisher", entity.Publisher, "empty" )
    addBodyTerm( &buffer, "datacite.creator", entity.Creator, "empty" )
    addBodyTerm( &buffer, "datacite.publicationyear", entity.PubYear, "empty" )
    addBodyTerm( &buffer, "datacite.resourcetype", entity.ResourceType, "Other" )
    s := buffer.String( )

    if debugHttp {
        fmt.Println("Payload:", s)
    }
    return s
}

func addBodyTerm( buffer * bytes.Buffer, term string, value string, defaultValue string ) {
    //fmt.Printf( "[%s] -> [%s]\n", term, value )

    if len( value ) != 0 {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, value ) )
    } else {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, defaultValue ) )
    }
}

func blankEntity( ) api.Entity {
    return api.Entity{ }
}

func statusIsOk( body string ) bool {
    //fmt.Println( "Response:", body )
    return( strings.Index( body, "success:" ) == 0 )
}
