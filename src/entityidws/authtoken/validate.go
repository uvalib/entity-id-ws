package authtoken

import (
    "fmt"
    "log"
    "time"
    "net/http"
    "github.com/parnurzeal/gorequest"
)

func Validate( endpoint string, activity string, token string ) bool {

    url := fmt.Sprintf( "%s/authorize/%s/%s/%s", endpoint, "entityidservice", activity, token )
    //log.Printf( "%s\n", url )

    start := time.Now( )
    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( 2 ) * time.Second ).
       End( )
    duration := time.Since( start )

    if errs != nil {
        log.Printf( "ERROR: token auth (%s) returns %s\n", url, errs )
        return false
    }

    defer resp.Body.Close( )

    log.Printf( "Token auth (%s) returns http %d in %s\n", url, resp.StatusCode, duration )
    return resp.StatusCode == http.StatusOK
}
