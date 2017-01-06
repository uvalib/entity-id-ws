package handlers

import (
    "net/http"
    "runtime"
)

func RuntimeInfo( w http.ResponseWriter, r *http.Request ) {

    ncpu := runtime.NumCPU()
    ngr := runtime.NumGoroutine( )
    m := &runtime.MemStats{ }
    runtime.ReadMemStats( m )
    encodeRuntimeResponse( w, http.StatusOK, ncpu, ngr, m.HeapObjects, m.Alloc )
}