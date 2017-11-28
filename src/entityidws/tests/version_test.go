package tests

import (
	"entityidws/client"
	"net/http"
	"testing"
)

//
// version tests
//

func TestVersionCheck(t *testing.T) {
	expected := http.StatusOK
	status, version := client.VersionCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(version) == 0 {
		t.Fatalf("Expected non-zero length version string\n")
	}
}

//
// end of file
//
