package handlers

import (
	"github.com/uvalib/entity-id-ws/entityidws/authtoken"
	"github.com/uvalib/entity-id-ws/entityidws/config"
	"github.com/uvalib/entity-id-ws/entityidws/idservice"
	"github.com/gorilla/mux"
	"net/http"
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
	if authtoken.Validate(config.Configuration.SharedSecret, token) == false {
		encodeStandardResponse(w, http.StatusForbidden)
		return
	}

	status := idservice.DeleteDoi(doi)
	encodeStandardResponse(w, status)
}

//
// end of file
//
