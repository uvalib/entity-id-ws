package main

import (
    "io/ioutil"
    "log"
    "testing"
    "entityidws/client"
    "gopkg.in/yaml.v2"
    "net/http"
    "strings"
    "entityidws/api"
)

type TestConfig struct {
    Endpoint  string
    Token     string
}

var cfg = loadConfig( )

var plausableDoi = "doi:10.5072/FK2QJ7DN9V"
var badDoi = "badness"
var goodShoulder = "doi:10.5072/FK2"
var badShoulder = "abc:/blablabla"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

//
// healthcheck tests
//

func TestHealthCheck( t *testing.T ) {
    expected := http.StatusOK
    status := client.HealthCheck( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// version tests
//

func TestVersionCheck( t *testing.T ) {
    expected := http.StatusOK
    status, version := client.VersionCheck( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }

    if len( version ) == 0 {
        t.Fatalf( "Expected non-zero length version string\n" )
    }
}

//
// DOI get tests
//

func TestGetHappyDay( t *testing.T ) {

    doi := createGoodDoi( )
    expected := http.StatusOK
    status, entity := client.Get( cfg.Endpoint, doi, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }

    if entity == nil {
        t.Fatalf( "Expected to find entity %v and did not\n", doi )
    }

    if emptyField( entity.Id ) ||
       emptyField( entity.Url ) ||
       emptyField( entity.Title ) ||
       emptyField( entity.Publisher ) ||
       emptyField( entity.Creator ) ||
       emptyField( entity.PubYear ) ||
       emptyField( entity.ResourceType ) {
        t.Fatalf( "Expected non-empty field but one is empty\n" )
    }
}

func TestGetEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.Get( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetBadId( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.Get( cfg.Endpoint, badDoi, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetEmptyToken( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.Get( cfg.Endpoint, plausableDoi, empty )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status, _ := client.Get( cfg.Endpoint, plausableDoi, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// DOI create tests
//

func TestCreateHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, entity := client.Create( cfg.Endpoint, goodShoulder, goodToken)
    if status != expected {
        t.Fatalf("Expected %v, got %v\n", expected, status)
    }

    if entity == nil {
        t.Fatalf("Expected to create entity successfully and did not\n" )
    }

    if emptyField( entity.Id ) {
        t.Fatalf( "Expected non-empty ID field but it is empty\n" )
    }
}

func TestCreateEmptyToken( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.Create( cfg.Endpoint, goodShoulder, empty )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestCreateBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status, _ := client.Create( cfg.Endpoint, goodShoulder, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// DOI update tests
//

func TestUpdateHappyDay( t *testing.T ) {

    doi := createGoodDoi( )
    entity := testEntity( )
    entity.Id = doi

    expected := http.StatusOK
    status := client.Update( cfg.Endpoint, entity, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestUpdateEmptyId( t *testing.T ) {
    entity := testEntity( )
    entity.Id = empty
    expected := http.StatusBadRequest
    status := client.Update( cfg.Endpoint, entity, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestUpdateBadId( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.Update( cfg.Endpoint, api.Entity{ Id: badDoi }, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestUpdateEmptyToken( t *testing.T ) {
    entity := testEntity( )
    entity.Id = plausableDoi
    expected := http.StatusBadRequest
    status := client.Update( cfg.Endpoint, entity, empty )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestUpdateBadToken( t *testing.T ) {
    entity := testEntity( )
    entity.Id = plausableDoi
    expected := http.StatusForbidden
    status := client.Update( cfg.Endpoint, entity, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// DOI delete tests
//

func TestDeleteEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.Delete( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteBadId( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.Delete( cfg.Endpoint, badDoi, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteEmptyToken( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.Delete( cfg.Endpoint, plausableDoi, empty )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status := client.Delete( cfg.Endpoint, plausableDoi, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// helpers
//

func emptyField( field string ) bool {
    return len( strings.TrimSpace( field ) ) == 0
}

func testEntity( ) api.Entity {
    return api.Entity{ Title: "my special title" }
}

func createGoodDoi( ) string {
    status, entity := client.Create( cfg.Endpoint, goodShoulder, goodToken )
    if status == http.StatusOK {
        return entity.Id
    }

    return ""
}

func loadConfig( ) TestConfig {

    data, err := ioutil.ReadFile( "service_test.yml" )
    if err != nil {
        log.Fatal( err )
    }

    var c TestConfig
    if err := yaml.Unmarshal( data, &c ); err != nil {
        log.Fatal( err )
    }

    log.Printf( "endpoint [%s]\n", c.Endpoint )
    log.Printf( "token    [%s]\n", c.Token )

    return c
}