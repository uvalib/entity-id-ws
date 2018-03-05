package tests

import (
	"entityidws/client"
	"net/http"
	"strings"
	"testing"
)

//
// metrics tests
//

func TestMetricsCheck(t *testing.T) {
	expected := http.StatusOK
	status, metrics := client.MetricsCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(metrics) == 0 {
		t.Fatalf("Expected non-empty metrics info\n")
	}

	if strings.Contains(metrics, "go_info") == false {
		t.Fatalf("Expected go_info value in metrics info\n")
	}
}

//
// end of file
//
