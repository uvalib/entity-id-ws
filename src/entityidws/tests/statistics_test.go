package tests

import (
   "entityidws/client"
   "net/http"
   "testing"
)

func TestStatistics(t *testing.T) {
   expected := http.StatusOK
   status, _ := client.Statistics(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   //if len( version ) == 0 {
   //    t.Fatalf( "Expected non-zero length version string\n" )
   //}
}

//
// end of file
//