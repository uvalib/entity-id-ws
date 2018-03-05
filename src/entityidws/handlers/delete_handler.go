package handlers

import (
	"entityidws/authtoken"
	"entityidws/config"
	"entityidws/idservice"
	"net/http"
	"github.com/gorilla/mux"
)

//
// IDDelete -- the delete by ID request handler
//
func IDDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	doi := vars["doi"]
	token := r.URL.Query().Get("auth")

	// validate inbound parameters
	if parameterOK(doi) == false || parameterOK(token) == false {
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, "delete", token, config.Configuration.Timeout) == false {
		encodeStandardResponse(w, http.StatusForbidden)
		return
	}

	status := idservice.DeleteDoi(doi)
	encodeStandardResponse(w, status)
}

//
// end of file
//
