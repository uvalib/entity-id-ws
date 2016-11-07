package ezid

import (
    "fmt"
    "strings"
    "bytes"
    "text/template"
    "entityidws/api"
    "entityidws/config"
    "entityidws/logger"
    "gopkg.in/xmlpath.v1"
    "html"
)

const PLACEHOLDER_TBA = "(:tba)"

//
// log the contents of an entity record
//
func logEntity( entity api.Entity ) {

    if config.Configuration.Debug {
        fmt.Println( "Id:", entity.Id )
        fmt.Println( "Url:", entity.Url )
        fmt.Println( "Title:", entity.Title )
        fmt.Println( "Publisher:", entity.Publisher )
        fmt.Println( "CreatorFirstName:", entity.CreatorFirstName )
        fmt.Println( "CreatorLastName:", entity.CreatorLastName )
        fmt.Println( "CreatorDepartment:", entity.CreatorDepartment )
        fmt.Println( "CreatorInstitution:", entity.CreatorInstitution )
        fmt.Println( "PublicationDate:", entity.PublicationDate )
        fmt.Println( "PublicationMilestone:", entity.PublicationMilestone )
        fmt.Println( "ResourceType:", entity.ResourceType )
    }
}

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
                t := strings.Split( s, "," )
                if len( t ) > 0 {
                    entity.CreatorLastName = t[ 0 ]
                }
                if len( t ) > 1 {
                    entity.CreatorFirstName = t[ 1 ]
                }
            case "datacite.publicationyear":
                entity.PublicationDate = s
            case "datacite.resourcetype":
                entity.ResourceType = s
            case "crossref":
                // our payload is a CrossRef XML schema, process as appropriate
                extractCrossRefData( &entity, s )
            }
        }
    }
    return entity

}

//
// use the datacite schema/profile to encode the metadata into the request body
//
func makeDataciteBodyFromEntity( entity api.Entity, status string ) ( string, error ) {

    // parse the publication date
    YYYY, _, _ := splitDate( entity.PublicationDate )

    creator := fmt.Sprintf( "%s, %s", entity.CreatorLastName, entity.CreatorFirstName )

    var buffer bytes.Buffer
    //addBodyTerm( &buffer, "_crossref", "no", "" )
    //addBodyTerm( &buffer, "_profile", "datacite", "" )
    addBodyTerm( &buffer, "_status", status, "reserved" )
    addBodyTerm( &buffer, "_target", entity.Url, "https://virginia.edu" )
    addBodyTerm( &buffer, "datacite.title", entity.Title, "empty" )
    addBodyTerm( &buffer, "datacite.publisher", entity.Publisher, "empty" )
    addBodyTerm( &buffer, "datacite.creator", creator, "empty" )
    addBodyTerm( &buffer, "datacite.publicationyear", YYYY, "empty" )
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
        entity.Id = PLACEHOLDER_TBA
        entity.Url = PLACEHOLDER_TBA
    }

    // parse the publication date
    YYYY, MM, DD := splitDate( entity.PublicationDate )

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
    } { htmlEncode( entity.CreatorFirstName ),
        htmlEncode( entity.CreatorLastName ),
        htmlEncode( entity.CreatorInstitution ),
        htmlEncode( entity.Title ),
        YYYY,
        MM,
        DD,
        htmlEncode( entity.CreatorDepartment ),
        htmlEncode( entity.PublicationMilestone ),
        entity.Id,
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

//
// extract data from the CrossRef schema
//
func extractCrossRefData( e * api.Entity, xref string ) {

    reader := strings.NewReader( xref )
    xmlroot, err := xmlpath.Parse( reader )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: parsing response XML: %s", err ) )
        return
    }

    //
    // pull out the data from the XML schema
    //
    val := extractFromSchema( xmlroot, "/dissertation/doi_data/doi" )
    if val != PLACEHOLDER_TBA {
        e.Id = val
    }
    val = extractFromSchema( xmlroot, "/dissertation/doi_data/resource" )
    if val != PLACEHOLDER_TBA {
        e.Url = val
    }
    e.Title = extractFromSchema( xmlroot, "/dissertation/titles/title" )
//    e.Publisher = extractFromSchema( xmlroot, "/dissertation/xxx" )
    e.CreatorFirstName = extractFromSchema( xmlroot, "/dissertation/person_name/given_name" )
    e.CreatorLastName = extractFromSchema( xmlroot, "/dissertation/person_name/surname" )
    e.CreatorDepartment = extractFromSchema( xmlroot, "/dissertation/institution/institution_department" )
    e.CreatorInstitution = extractFromSchema( xmlroot, "/dissertation/person_name/affiliation" )

    e.PublicationDate = extractFromSchema( xmlroot, "/dissertation/approval_date/year" )
    MM := extractFromSchema( xmlroot, "/dissertation/approval_date/month" )
    if len( MM ) > 0 {
        e.PublicationDate = fmt.Sprintf( "%s-%s", e.PublicationDate, MM )
    }
    DD := extractFromSchema( xmlroot, "/dissertation/approval_date/day" )
    if len( DD ) > 0 {
        e.PublicationDate = fmt.Sprintf( "%s-%s", e.PublicationDate, DD )
    }

    e.PublicationMilestone = extractFromSchema( xmlroot, "/dissertation/degree" )
}

func addBodyTerm( buffer * bytes.Buffer, term string, value string, defaultValue string ) {
    if len( value ) != 0 {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, specialEncode( value ) ) )
    } else {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, specialEncode( defaultValue ) ) )
    }
}

//
// the EZID service requires that embedded newlines and carriage returns be percent encoded.
//
func specialEncode( value string ) string {

    // EZID structural element encoding
    value = strings.Replace( value, "\n", "%0A", -1 )
    value = strings.Replace( value, "\r", "%0B", -1 )
    return value
}

//
// when including content embedded in XML, we should HTML encode it.
//
func htmlEncode( value string ) string {
    // HTML encoding
    return html.EscapeString( value )
}

//
// create a blank entity
//
func blankEntity( ) api.Entity {
    return api.Entity{ }
}

//
// extract from schema
//
func extractFromSchema( xmlroot * xmlpath.Node, xpath string ) string {
    path := xmlpath.MustCompile( xpath )
    if value, ok := path.String( xmlroot ); ok {
        return value
    }

    return ""
}

//
// check for an OK response status
//
func statusIsOk( body string ) bool {
    //fmt.Println( "Response:", body )
    return( strings.Index( body, "success:" ) == 0 )
}

//
// Split a date in the form YYYY-MM-DD into its components
//
func splitDate( date string ) ( string, string, string ) {
    tokens := strings.Split( date, "-" )
    var YYYY, MM, DD string
    if len( tokens ) > 0 {
        YYYY = tokens[ 0 ]
    }

    if len( tokens ) > 1 {
        MM = tokens[ 1 ]
    }

    if len( tokens ) > 2 {
        DD = tokens[ 2 ]
    }
    return YYYY, MM, DD
}
