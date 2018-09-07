package handlers

import (
	"net/http"
)

//
// FavIconHandler - get favicon handler
//
func FavIconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

//
// end of file
//
