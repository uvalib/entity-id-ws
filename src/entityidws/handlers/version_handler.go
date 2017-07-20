package handlers

import (
	"net/http"
)

func VersionInfo(w http.ResponseWriter, r *http.Request) {
	encodeVersionResponse(w, http.StatusOK, Version())
}
