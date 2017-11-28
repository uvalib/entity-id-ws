package tests

import (
   "entityidws/client"
   "net/http"
   "testing"
)

func TestCreateCrossRef(t *testing.T) {
   expected := http.StatusOK

   status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), goodToken)
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

func TestCreateDataCite(t *testing.T) {
   expected := http.StatusOK

   status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(dataciteSchema), goodToken)
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

   status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(badSchema), goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestCreateEmptyToken(t *testing.T) {
   expected := http.StatusBadRequest

   status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), empty)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestCreateBadToken(t *testing.T) {
   expected := http.StatusForbidden

   status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//