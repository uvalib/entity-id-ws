package tests

import (
	"entityidws/client"
	"net/http"
	"testing"
)

//
// runtime tests
//

func TestRuntimeCheck(t *testing.T) {
	expected := http.StatusOK
	status, runtime := client.RuntimeCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if runtime == nil {
		t.Fatalf("Expected non-nil runtime info\n")
	}

	if len(runtime.Version) == 0 ||
		runtime.AllocatedMemory == 0 ||
		runtime.CPUCount == 0 ||
		runtime.GoRoutineCount == 0 ||
		runtime.ObjectCount == 0 {
		t.Fatalf("Expected non-zero value in runtime info but one is zero\n")
	}
}

//
// end of file
//
