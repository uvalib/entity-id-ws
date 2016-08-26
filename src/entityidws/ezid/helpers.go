package ezid

import (
    "fmt"
    "strings"
    "bytes"
    "text/template"
    "entityidws/api"
    "entityidws/config"
    "entityidws/logger"
)

//
// the response body consists of a set of CR separated lines containing tokens separated by
// a colon character
//
func makeEntityFromBody( body string ) api.Entity {

    if config.Configuration.Debug {
        fmt.Println("Response:", body)
    }

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
func makeDataciteBodyFromEntity( entity api.Entity, status string ) ( string, error ) {

    var buffer bytes.Buffer
    //addBodyTerm( &buffer, "_crossref", "no", "" )
    //addBodyTerm( &buffer, "_profile", "datacite", "" )
    addBodyTerm( &buffer, "_status", status, "reserved" )
    addBodyTerm( &buffer, "_target", entity.Url, "https://virginia.edu" )
    addBodyTerm( &buffer, "datacite.title", entity.Title, "empty" )
    addBodyTerm( &buffer, "datacite.publisher", entity.Publisher, "empty" )
    addBodyTerm( &buffer, "datacite.creator", entity.Creator, "empty" )
    addBodyTerm( &buffer, "datacite.publicationyear", entity.PubYear, "empty" )
    addBodyTerm( &buffer, "datacite.resourcetype", entity.ResourceType, "Other" )
    s := buffer.String( )

    if config.Configuration.Debug {
        fmt.Println( "Payload:", s )
    }
    return s, nil
}

//
// use the crossref schema/profile to encode the metadata into the request body
//
func makeCrossRefBodyFromEntity( entity api.Entity, status string ) ( string, error ) {

    // create the XML payload
    xref, err := createCrossRefSchema( entity, status )
    if err != nil {
        return "", err
    }

    var buffer bytes.Buffer
    addBodyTerm( &buffer, "_crossref", "yes", "" )
    addBodyTerm( &buffer, "_profile", "crossref", "" )
    addBodyTerm( &buffer, "_status", status, "reserved" )
    addBodyTerm( &buffer, "_target", entity.Url, "https://virginia.edu" )
    addBodyTerm( &buffer, "crossref", xref, "" )
    s := buffer.String( )

    if config.Configuration.Debug {
        fmt.Println( "Payload:", s )
    }
    return s, nil
}

//
// use the crossref template to encode the metadata
//
func createCrossRefSchema( entity api.Entity, status string ) ( string, error ) {

    t, err := template.ParseFiles( "data/crossref-template.xml" )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: template parse error: %s", err ) )
        return "", err
    }

    // add placeholders if we are reserving a DOI
    if status == STATUS_RESERVED {
        //entity.Doi = "(:tba)"
        entity.Url = "(:tba)"
    }

    // create our template data structure
    data := struct {
        FirstName   string
        LastName    string
        Institution string
        Title       string
        PubYear     string
        PubMonth    string
        PubDay      string
        Department  string
        Degree      string
        Identifier  string
        PublicUrl   string
    } { "Dave",
        "Goldstein",
        "UVA",
        entity.Title,
        entity.PubYear,
        "MM",
        "DD",
        "Department",
        "PHD",
        "(:tba)",
        entity.Url }

    var buffer bytes.Buffer
    err = t.Execute( &buffer, data )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: template execute error: %s", err ) )
        return "", err
    }

    s := buffer.String( )
    if config.Configuration.Debug {
        fmt.Printf( "XML:\n%s\n", s )
    }
    return s, nil
}

func addBodyTerm( buffer * bytes.Buffer, term string, value string, defaultValue string ) {
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

//
// check for an OK response status
//
func statusIsOk( body string ) bool {
    //fmt.Println( "Response:", body )
    return( strings.Index( body, "success:" ) == 0 )
}
