package tests

import (
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"net/http"
	"testing"
)

/*
func TestDeleteCrossRef(t *testing.T) {
	expected := http.StatusOK
	doi := createGoodDoi(crossrefSchema, t)
	status := client.Delete(cfg.Endpoint, doi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}
*/

func TestDeleteDataCite(t *testing.T) {
	expected := http.StatusOK
	doi := createGoodDoi(dataciteSchema, t)
	status := client.Delete(cfg.Endpoint, doi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Delete(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteBadId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Delete(cfg.Endpoint, badDoi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Delete(cfg.Endpoint, plausableDoi, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status := client.Delete(cfg.Endpoint, plausableDoi, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
