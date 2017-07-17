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
    "errors"
    //"path"
)

const PLACEHOLDER_TBA = "(:tba)"
const CROSSREF_SCHEMA = "crossref"
const DATACITE_SCHEMA = "datacite"

//
// log the contents of a request record
//
func logRequest( request api.Request) {

    if config.Configuration.Debug {
        fmt.Println( "Schema:", request.Schema )
        fmt.Println( "Id:", request.Id )

        if request.Schema == CROSSREF_SCHEMA {
            logCrossRefRequest( request.CrossRef )
        }
        if request.Schema == DATACITE_SCHEMA {
            logDataCiteRequest( request.DataCite )
        }
    }
}

//
// log the contents of a crossref schema request
//
func logCrossRefRequest( request api.CrossRefSchema ) {

    fmt.Println( "Url:", request.Url )
    fmt.Println( "Title:", request.Title )
    fmt.Println( "Publisher:", request.Publisher )
    fmt.Println( "CreatorFirstName:", request.CreatorFirstName )
    fmt.Println( "CreatorLastName:", request.CreatorLastName )
    fmt.Println( "CreatorDepartment:", request.CreatorDepartment )
    fmt.Println( "CreatorInstitution:", request.CreatorInstitution )
    fmt.Println( "PublicationDate:", request.PublicationDate )
    fmt.Println( "PublicationMilestone:", request.PublicationMilestone )
    fmt.Println( "ResourceType:", request.ResourceType )
}

//
// log the contents of a datacite schema request
//
func logDataCiteRequest( request api.DataCiteSchema ) {

    fmt.Println( "Url:", request.Url )
    fmt.Println( "Title:", request.Title )
    fmt.Println( "Abstract:", request.Abstract )
    fmt.Println( "Creators:", request.Creators )
    fmt.Println( "Contributors:", request.Contributors )
    fmt.Println( "Rights:", request.Rights )
    fmt.Println( "Keywords:", request.Keywords )
    fmt.Println( "Sponsors:", request.Sponsors )
    fmt.Println( "Publisher:", request.Publisher )
    fmt.Println( "PublicationDate:", request.PublicationDate )
    fmt.Println( "GeneralType:", request.GeneralType )
    fmt.Println( "ResourceType:", request.ResourceType )
}

//
// the response body consists of a set of CR separated lines containing tokens separated by
// a colon character
//
func makeEntityFromBody( body string ) api.Request {

    if config.Configuration.Debug {
        fmt.Println("Response:", body)
    }

    response := blankResponse( )
    split := strings.Split( body, "\n" )
    for i := range split {
        tokens := strings.SplitN( split[ i ], ":", 2 )
        if len( tokens ) == 2 {
            s := strings.TrimSpace( tokens[ 1 ] )
            switch tokens[ 0 ] {
            case "success":
                response.Id = strings.TrimSpace( strings.Split( s, "|" )[ 0 ] )
            case "_profile":
                response.Schema = s
            case DATACITE_SCHEMA:
                // our payload is a DataCite XML schema, process as appropriate
                response.Schema = DATACITE_SCHEMA
                extractDataCitePayload( &response, s )
            case CROSSREF_SCHEMA:
                // our payload is a CrossRef XML schema, process as appropriate
                response.Schema = CROSSREF_SCHEMA
                extractCrossRefPayload( &response, s )
            }
        }
    }
    return response

}

//
// encode the data into the request body
//
func makeBodyFromRequest( request api.Request, status string ) ( string, error ) {

    var body string
    var err error

    // check the schema type and build the body as appropriate
    switch request.Schema {
    case CROSSREF_SCHEMA:
        body, err = makeCrossRefBodyFromEntity( request, status )
    case DATACITE_SCHEMA:
        body, err = makeDataCiteBodyFromEntity( request, status )
    default:
        return "", errors.New( fmt.Sprintf( "unregognized schema name: %s", request.Schema ) )
    }

    if err != nil {
        return "", err
    }

    if config.Configuration.Debug {
        fmt.Println( "Payload:", body )
    }

    return body, nil
}

//
// use the datacite schema/profile to encode the metadata into the request body
//
func makeDataCiteBodyFromEntity( request api.Request, status string ) ( string, error ) {

    // create the XML payload
    xml, err := createDataCiteSchema( request, status )
    if err != nil {
        return "", err
    }

    var buffer bytes.Buffer
    addBodyTerm( &buffer, "_profile", "datacite", "" )
    addBodyTerm( &buffer, "_status", status, "reserved" )
    addBodyTerm( &buffer, "_target", request.DataCite.Url, "https://virginia.edu" )
    addBodyTerm( &buffer, "datacite", xml, "" )
    s := buffer.String( )
    return s, nil
}

//
// use the crossref schema/profile to encode the metadata into the request body
//
func makeCrossRefBodyFromEntity( request api.Request, status string ) ( string, error ) {

    // create the XML payload
    xml, err := createCrossRefSchema( request, status )
    if err != nil {
        return "", err
    }

    var buffer bytes.Buffer
    addBodyTerm( &buffer, "_crossref", "yes", "" )
    addBodyTerm( &buffer, "_profile", "crossref", "" )
    addBodyTerm( &buffer, "_status", status, "reserved" )
    addBodyTerm( &buffer, "_target", request.CrossRef.Url, "https://virginia.edu" )
    addBodyTerm( &buffer, "crossref", xml, "" )
    s := buffer.String( )
    return s, nil
}

//
// use the datacite template to encode the metadata
//
func createDataCiteSchema( request api.Request, status string ) ( string, error ) {

    t, err := template.ParseFiles("data/datacite-template.xml")
    if err != nil {
        logger.Log(fmt.Sprintf("ERROR: template parse error: %s", err))
        return "", err
    }

    // add placeholders if we are reserving a DOI
    if status == STATUS_RESERVED {
        request.Id = PLACEHOLDER_TBA
        request.DataCite.Url = PLACEHOLDER_TBA
    }

    // parse the publication date
    YYYY, _, _ := splitDate( request.DataCite.PublicationDate )

    // create our template data structure
    data := struct {
        Identifier       string
        Title            string
        Abstract         string
        Creators      [] api.Person
        Contributors  [] api.Person
        Rights           string
        Publisher        string
        PublicationDate  string
        PublicationYear  string
        Keywords      [] string
        Sponsors      [] string
        GeneralType      string
        ResourceType     string

    } {
        request.Id,
        htmlEncodeString( orUnavailable( request.DataCite.Title ) ),
        htmlEncodeString( orUnavailable( request.DataCite.Abstract ) ),
        htmlEncodePersonArray( api.SortPeople( request.DataCite.Creators ) ),
        htmlEncodePersonArray( api.SortPeople( request.DataCite.Contributors ) ),
        htmlEncodeString( orUnavailable( request.DataCite.Rights ) ),
        htmlEncodeString( orUnavailable( request.DataCite.Publisher ) ),
        request.DataCite.PublicationDate,
        YYYY,
        htmlEncodeStringArray( request.DataCite.Keywords ),
        htmlEncodeStringArray( request.DataCite.Sponsors ),
        request.DataCite.GeneralType,
        request.DataCite.ResourceType,
    }

    var buffer bytes.Buffer
    err = t.Execute( &buffer, data )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: template execute error: %s", err ) )
        return "", err
    }

    s := buffer.String( )
    s = filterUnacceptableCharacters( s )

    if config.Configuration.Debug {
        fmt.Printf( "XML:\n%s\n", s )
    }
    return s, nil
}

//
// use the crossref template to encode the metadata
//
func createCrossRefSchema( request api.Request, status string ) ( string, error ) {

    t, err := template.ParseFiles( "data/crossref-template.xml" )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: template parse error: %s", err ) )
        return "", err
    }

    // add placeholders if we are reserving a DOI
    if status == STATUS_RESERVED {
        request.Id = PLACEHOLDER_TBA
        request.CrossRef.Url = PLACEHOLDER_TBA
    }

    // parse the publication date
    YYYY, MM, DD := splitDate( request.CrossRef.PublicationDate )

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
    } { htmlEncodeString( request.CrossRef.CreatorFirstName ),
        htmlEncodeString( request.CrossRef.CreatorLastName ),
        htmlEncodeString( request.CrossRef.CreatorInstitution ),
        htmlEncodeString( request.CrossRef.Title ),
        YYYY,
        MM,
        DD,
        htmlEncodeString( request.CrossRef.CreatorDepartment ),
        htmlEncodeString( request.CrossRef.PublicationMilestone ),
        request.Id,
        request.CrossRef.Url }

    var buffer bytes.Buffer
    err = t.Execute( &buffer, data )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: template execute error: %s", err ) )
        return "", err
    }

    s := buffer.String( )
    s = filterUnacceptableCharacters( s )

    if config.Configuration.Debug {
        fmt.Printf( "XML:\n%s\n", s )
    }
    return s, nil
}

//
// extract data from the DataCite schema
//
func extractDataCitePayload( payload * api.Request, xml string ) {

    reader := strings.NewReader(xml)
    xmlroot, err := xmlpath.Parse(reader)
    if err != nil {
        logger.Log(fmt.Sprintf("ERROR: parsing response XML: %s", err))
        return
    }

    //
    // pull out the data from the XML schema
    //
    val := extractStringFromSchema( xmlroot, "/resource/identifier" )
    if val != PLACEHOLDER_TBA {
        payload.Id = val
    }

    payload.DataCite.Title = extractStringFromSchema( xmlroot, "/resource/titles/title" )
    payload.DataCite.Abstract = extractStringFromSchema( xmlroot, "/resource/descriptions/description" )

    payload.DataCite.Creators = extractPersonListFromSchema( xmlroot, "/resource/creators/creator" )
    payload.DataCite.Contributors = extractPersonListFromSchema( xmlroot, "/resource/contributors/contributor" )
    payload.DataCite.Rights = extractStringFromSchema( xmlroot, "/resource/rightsList/rights" )
    payload.DataCite.Keywords = extractStringListFromSchema( xmlroot, "/resource/subjects/subject" )
    payload.DataCite.Sponsors = extractStringListFromSchema( xmlroot, "/resource/fundingReferences/fundingReference/funderName" )
    payload.DataCite.Publisher = extractStringFromSchema( xmlroot, "/resource/publisher" )
    payload.DataCite.PublicationDate = extractStringFromSchema( xmlroot, "/resource/dates/date" )
    payload.DataCite.GeneralType = extractStringFromSchema( xmlroot, "/resource/resourceType/@resourceTypeGeneral" )
    payload.DataCite.ResourceType = extractStringFromSchema( xmlroot, "/resource/resourceType" )
}

//
// extract data from the CrossRef schema
//
func extractCrossRefPayload( payload * api.Request, xml string ) {

    reader := strings.NewReader( xml )
    xmlroot, err := xmlpath.Parse( reader )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: parsing response XML: %s", err ) )
        return
    }

    //
    // pull out the data from the XML schema
    //
    val := extractStringFromSchema( xmlroot, "/dissertation/doi_data/doi" )
    if val != PLACEHOLDER_TBA {
        payload.Id = val
    }
    val = extractStringFromSchema( xmlroot, "/dissertation/doi_data/resource" )
    if val != PLACEHOLDER_TBA {
        payload.CrossRef.Url = val
    }
    payload.CrossRef.Title = extractStringFromSchema( xmlroot, "/dissertation/titles/title" )
    payload.CrossRef.CreatorFirstName = extractStringFromSchema( xmlroot, "/dissertation/person_name/given_name" )
    payload.CrossRef.CreatorLastName = extractStringFromSchema( xmlroot, "/dissertation/person_name/surname" )
    payload.CrossRef.CreatorDepartment = extractStringFromSchema( xmlroot, "/dissertation/institution/institution_department" )
    payload.CrossRef.CreatorInstitution = extractStringFromSchema( xmlroot, "/dissertation/person_name/affiliation" )

    payload.CrossRef.PublicationDate = extractStringFromSchema( xmlroot, "/dissertation/approval_date/year" )
    MM := extractStringFromSchema( xmlroot, "/dissertation/approval_date/month" )
    if len( MM ) > 0 {
        payload.CrossRef.PublicationDate = fmt.Sprintf( "%s-%s", payload.CrossRef.PublicationDate, MM )
    }
    DD := extractStringFromSchema( xmlroot, "/dissertation/approval_date/day" )
    if len( DD ) > 0 {
        payload.CrossRef.PublicationDate = fmt.Sprintf( "%s-%s", payload.CrossRef.PublicationDate, DD )
    }

    payload.CrossRef.PublicationMilestone = extractStringFromSchema( xmlroot, "/dissertation/degree" )
}

func addBodyTerm( buffer * bytes.Buffer, term string, value string, defaultValue string ) {
    if len( value ) != 0 {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, specialEncode( value ) ) )
    } else {
        buffer.WriteString( fmt.Sprintf( "%s: %s\n", term, specialEncode( defaultValue ) ) )
    }
}

//
// the EZID service requires that embedded newlines be percent encoded.
//
func specialEncode( value string ) string {

    value = strings.Replace( value, "\n", "%0A", -1 )
    return value
}

//
// certain characters are unacceptable in the payload.
//
func filterUnacceptableCharacters( value string ) string {

    value = strings.Replace( value, "\r", "", -1 )
    return value
}

//
// when including content embedded in XML, we should HTML encode it.
//
func htmlEncodePersonArray( array [] api.Person ) [] api.Person {

    encoded := make([] api.Person, len( array ), len( array ) )
    for ix, value := range array {

        p := api.Person{
            Index:       value.Index,
            FirstName:   htmlEncodeString( value.FirstName ),
            LastName:    htmlEncodeString( value.LastName ),
            Department:  htmlEncodeString( value.Department ),
            Institution: htmlEncodeString( value.Institution ),
        }
        encoded[ ix ] = p
    }
    return encoded
}

func htmlEncodeStringArray( array [] string ) [] string {

    encoded := make([] string, len( array ), len( array ) )
    for ix, value := range array {
        encoded[ ix ] = htmlEncodeString( value )
    }
    return encoded
}

func htmlEncodeString( value string ) string {
    // HTML encoding
    encoded := html.EscapeString( value )

    // encode percent characters
    encoded = strings.Replace( encoded, "%", "%25", -1 )
    return encoded
}

//
// create a blank entity
//
func blankResponse( ) api.Request {
    return api.Request{ }
}

//
// extract a string from the schema
//
func extractStringFromSchema( xmlroot * xmlpath.Node, xpath string ) string {
    path := xmlpath.MustCompile( xpath )
    if value, ok := path.String( xmlroot ); ok {
        return value
    }

    return ""
}

//
// extract a list of strings from the schema
//
func extractStringListFromSchema( xmlroot * xmlpath.Node, xpath string ) [] string {
    var result []string
    path := xmlpath.MustCompile( xpath )
    iter := path.Iter( xmlroot )
    for iter.Next() {
        result = append( result, iter.Node( ).String( ) )
    }
    return result
}

//
// extract a list of strings from the schema
//
func extractPersonListFromSchema( xmlroot * xmlpath.Node, xpath string ) [] api.Person {
    var result []api.Person
    found := 0

    path := xmlpath.MustCompile( xpath )
    iter := path.Iter( xmlroot )
    for iter.Next() {
        found += 1
    }
    if found > 0 {
        for i := 0; i < found; i++ {
            gn := extractStringFromSchema( xmlroot, fmt.Sprintf("%s[%d]/givenName", xpath, i + 1 ) )
            fn := extractStringFromSchema( xmlroot, fmt.Sprintf("%s[%d]/familyName", xpath, i + 1 ) )
            af := extractStringFromSchema( xmlroot, fmt.Sprintf("%s[%d]/affiliation", xpath, i + 1 ) )
            p := api.Person{ Index: i, FirstName: gn, LastName: fn, Institution: af }
            result = append( result, p )
        }
    }

    return result
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

//
// helper to return a standard code for a blank field. In this case, the unavailable one
//
func orUnavailable( field string ) string {

    if len( field ) != 0 {
        return( field )
    }
    return( "(:unav)" )
}
