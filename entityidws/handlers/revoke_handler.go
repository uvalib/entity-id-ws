package handlers

import (
	"github.com/uvalib/entity-id-ws/entityidws/authtoken"
	"github.com/uvalib/entity-id-ws/entityidws/config"
	"github.com/uvalib/entity-id-ws/entityidws/idservice"
	"net/http"

	"github.com/gorilla/mux"
)

//
// IDRevoke -- the revoke by ID request handler
//
func IDRevoke(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	doi := vars["doi"]
	token := r.URL.Query().Get("auth")

	// validate inbound parameters
	if parameterOK(doi) == false || parameterOK(token) == false {
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, "delete", token, config.Configuration.ServiceTimeout) == false {
		encodeStandardResponse(w, http.StatusForbidden)
		return
	}

	// get the existing metadata
	entity, status := idservice.GetDoi(doi)
	if status == http.StatusOK {

		// update the status
		entity.ID = doi
		status = idservice.UpdateDoi(entity, idservice.StatusUnavailable)
	}

	encodeStandardResponse(w, status)
}

//
// end of file
//
