package ezid

import (
    "fmt"
    "strings"
    "bytes"
    "entityidws/api"
)

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

//
// use the datacite schema/profile to encode the metadata into the request body
//
func makeDataciteBodyFromEntity( entity api.Entity ) string {

    var buffer bytes.Buffer
    //addBodyTerm( &buffer, "_crossref", "no", "" )
    //addBodyTerm( &buffer, "_profile", "datacite", "" )
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
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, specialEncode( value ) ) )
    } else {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, specialEncode( defaultValue ) ) )
    }
}

//
// the EZID service requires that embedded newlines and carriage returns be percent encoded
//
func specialEncode( value string ) string {

    value = strings.Replace( value, "\n", "%0A", -1 )
    value = strings.Replace( value, "\r", "%0B", -1 )
    return value
}

func blankEntity( ) api.Entity {
    return api.Entity{ }
}

func statusIsOk( body string ) bool {
    //fmt.Println( "Response:", body )
    return( strings.Index( body, "success:" ) == 0 )
}