package tests

import (
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"net/http"
	"testing"
)

/*
func TestUpdateCrossRef(t *testing.T) {

	doi := createGoodDoi(crossrefSchema, t)
	entity := createTestRequest(crossrefSchema)
	entity.ID = doi

	expected := http.StatusOK
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}
*/

func TestUpdateDataCite(t *testing.T) {

	doi := createGoodDoi(dataciteSchema, t)
	entity := createTestRequest(dataciteSchema)
	entity.ID = doi

	expected := http.StatusOK
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateBadSchema(t *testing.T) {

	doi := createGoodDoi(dataciteSchema, t)
	entity := createTestRequest(badSchema)
	entity.ID = doi

	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateEmptyId(t *testing.T) {
	entity := createTestRequest(dataciteSchema)
	entity.ID = empty
	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateBadId(t *testing.T) {
	entity := createTestRequest(dataciteSchema)
	entity.ID = badDoi
	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateEmptyToken(t *testing.T) {
	entity := createTestRequest(dataciteSchema)
	entity.ID = plausableDoi
	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateBadToken(t *testing.T) {
	entity := createTestRequest(dataciteSchema)
	entity.ID = plausableDoi
	expected := http.StatusForbidden
	status := client.Update(cfg.Endpoint, entity, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
