package authtoken

import (
    "fmt"
    "time"
    "net/http"
    "github.com/parnurzeal/gorequest"
    "entityidws/logger"
)

func Validate( endpoint string, activity string, token string ) bool {

    url := fmt.Sprintf( "%s/authorize/%s/%s/%s", endpoint, "entityidservice", activity, token )
    //log.Printf( "%s\n", url )

    start := time.Now( )
    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )
    duration := time.Since( start )

    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: token auth (%s) returns %s in %s\n", url, errs, duration ) )
        return false
    }

    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Token auth (%s) returns http %d in %s\n", url, resp.StatusCode, duration ) )
    return resp.StatusCode == http.StatusOK
}
