package handlers

import (
	"entityidws/authtoken"
	"entityidws/config"
	"entityidws/ezid"
	"github.com/gorilla/mux"
	"net/http"
)

//
// IDRevoke -- the revoke by ID request handler
//
func IDRevoke(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	doi := vars["doi"]
	token := r.URL.Query().Get("auth")

	// update the statistics
	Statistics.RequestCount++
	Statistics.RevokeCount++

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

	// get the existing metadata
	entity, status := ezid.GetDoi(doi)
	if status == http.StatusOK {

		// update the status
		entity.ID = doi
		status = ezid.UpdateDoi(entity, ezid.StatusUnavailable)
	}

	encodeStandardResponse(w, status)
}

//
// end of file
//
