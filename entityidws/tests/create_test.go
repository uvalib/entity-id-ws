package tests

import (
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"net/http"
	"testing"
)

/*
func TestCreateCrossRef(t *testing.T) {
	expected := http.StatusOK

	status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if entity == nil {
		t.Fatalf("Expected to create entity successfully and did not\n")
	}

	if emptyField(entity.ID) {
		t.Fatalf("Expected non-empty ID field but it is empty\n")
	}
}
*/

func TestCreateDataCite(t *testing.T) {
	expected := http.StatusOK

	status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(dataciteSchema), goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if entity == nil {
		t.Fatalf("Expected to create entity successfully and did not\n")
	}

	if emptyField(entity.ID) {
		t.Fatalf("Expected non-empty ID field but it is empty\n")
	}
}

func TestCreateBadSchema(t *testing.T) {
	expected := http.StatusBadRequest

	status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(badSchema), goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestCreateEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest

	status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(dataciteSchema), empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestCreateBadToken(t *testing.T) {
	expected := http.StatusForbidden

	status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(dataciteSchema), badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
