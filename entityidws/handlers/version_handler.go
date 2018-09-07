package handlers

import (
	"net/http"
)

//
// VersionInfo - get version handler
//
func VersionInfo(w http.ResponseWriter, r *http.Request) {
	encodeVersionResponse(w, http.StatusOK, Version())
}

//
// end of file
//
