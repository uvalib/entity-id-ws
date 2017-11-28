package tests

import (
   "entityidws/client"
   "net/http"
   "testing"
)

func TestRevokeCrossRef(t *testing.T) {

   expected := http.StatusOK
   doi := createGoodDoi(crossrefSchema, t)
   entity := createTestRequest(crossrefSchema)
   entity.ID = doi

   status := client.Update(cfg.Endpoint, entity, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   status = client.Revoke(cfg.Endpoint, entity.ID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestRevokeDataSite(t *testing.T) {

   expected := http.StatusOK
   doi := createGoodDoi(dataciteSchema, t)
   entity := createTestRequest(dataciteSchema)
   entity.ID = doi

   status := client.Update(cfg.Endpoint, entity, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   status = client.Revoke(cfg.Endpoint, entity.ID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestRevokeBadSchema(t *testing.T) {

   expected := http.StatusBadRequest
   doi := createGoodDoi(crossrefSchema, t)
   entity := createTestRequest(badSchema)
   entity.ID = doi

   status := client.Update(cfg.Endpoint, entity, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   status = client.Revoke(cfg.Endpoint, entity.ID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestRevokeEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.Revoke(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestRevokeBadId(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.Revoke(cfg.Endpoint, badDoi, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestRevokeEmptyToken(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.Revoke(cfg.Endpoint, plausableDoi, empty)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestRevokeBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status := client.Revoke(cfg.Endpoint, plausableDoi, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//