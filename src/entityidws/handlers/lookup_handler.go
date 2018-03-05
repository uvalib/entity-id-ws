package handlers

import (
	"entityidws/authtoken"
	"entityidws/config"
	"entityidws/idservice"
	"net/http"

	"github.com/gorilla/mux"
)

//
// IDLookup -- the lookup by id request handler
//
func IDLookup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	doi := vars["doi"]
	token := r.URL.Query().Get("auth")

	// validate inbound parameters
	if parameterOK(doi) == false || parameterOK(token) == false {
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, "lookup", token, config.Configuration.Timeout) == false {
		encodeStandardResponse(w, http.StatusForbidden)
		return
	}

	entity, status := idservice.GetDoi(doi)
	encodeDetailsResponse(w, status, entity)
}

//
// end of file
//
