package main

import (
    "fmt"
    "time"
    "net/http"
    "strings"
    "github.com/parnurzeal/gorequest"
)

//
// get entity details when provided a DOI
//
func GetByDoi( doi string ) ( Entity, int ) {

    url := fmt.Sprintf( "%s/id/%s", config.EzidServiceUrl, doi )
    //fmt.Println( "URL:", url )

    _, body, errs := gorequest.New( ).
       Get( url  ).
       Timeout( time.Duration( config.EzidServiceTimeout ) * time.Second ).
       End( )

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
    return makeEntity( body ), http.StatusOK
}

func CreateByDoi( entity Entity ) ( Entity, int ) {
    // all good...
    return blankEntity( ), http.StatusOK
}

func UpdateByDoi( entity Entity ) ( Entity, int ) {
    // all good...
    return blankEntity( ), http.StatusOK
}

func DeleteByDoi( doi string ) int {

    // all good...
    return http.StatusOK
}

//
// the response body consists of a set of CR separated lines containing
func makeEntity( body string ) Entity {

    //fmt.Println( "Response:", body )

    entity := blankEntity( )
    split := strings.Split( body, "\n" )
    for i := range split {
        tokens := strings.SplitN( split[i], ":", 2 )
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

func blankEntity( ) Entity {
   return Entity{ }
}

func statusIsOk( body string ) bool {
   return( strings.Index( body, "success:" ) == 0 )
}
