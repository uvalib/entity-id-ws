package tests

import (
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"net/http"
	"testing"
)

/*
func TestGetCrossRef(t *testing.T) {

	doi := createGoodDoi(crossrefSchema, t)
	expected := http.StatusOK
	status, response := client.Get(cfg.Endpoint, doi, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if response == nil {
		t.Fatalf("Expected to find entity %v and did not\n", doi)
	}

	if response.Schema != crossrefSchema {
		t.Fatalf("Received unexpected schema in response\n")
	}
	if emptyField(response.ID) {
		t.Fatalf("Received blank ID in response\n")
	}
	verifyCrossRefSchema(response.CrossRef, t)
}
*/

func TestGetDataCite(t *testing.T) {

	doi := createGoodDoi(dataciteSchema, t)
	expected := http.StatusOK
	status, response := client.Get(cfg.Endpoint, doi, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if response == nil {
		t.Fatalf("Expected to find entity %v and did not\n", doi)
	}

	if response.Schema != dataciteSchema {
		t.Fatalf("Received unexpected schema in response\n")
	}
	if emptyField(response.ID) {
		t.Fatalf("Received blank ID in response\n")
	}
	verifyDataCiteSchema(response.DataCite, t)
}

func TestGetEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.Get(cfg.Endpoint, empty, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetBadId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.Get(cfg.Endpoint, badDoi, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.Get(cfg.Endpoint, plausableDoi, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _ := client.Get(cfg.Endpoint, plausableDoi, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
