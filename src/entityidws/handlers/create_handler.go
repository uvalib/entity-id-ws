package handlers

import (
	"encoding/json"
	"entityidws/api"
	"entityidws/authtoken"
	"entityidws/config"
	"entityidws/ezid"
	"entityidws/logger"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

//
// IDCreate -- the create by ID request handler
//
func IDCreate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	shoulder := vars["shoulder"]
	token := r.URL.Query().Get("auth")

	// update the statistics
	Statistics.RequestCount++
	Statistics.CreateCount++

	// validate inbound parameters
	if parameterOK(shoulder) == false || parameterOK(token) == false {
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, "create", token, config.Configuration.Timeout) == false {
		encodeStandardResponse(w, http.StatusForbidden)
		return
	}

	decoder := json.NewDecoder(r.Body)
	request := api.Request{}

	if err := decoder.Decode(&request); err != nil {
		logger.Log(fmt.Sprintf("ERROR: decoding request payload %s", err))
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	defer io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	entity, status := ezid.CreateDoi(shoulder, request, ezid.StatusReserved)
	encodeDetailsResponse(w, status, entity)
}

//
// end of file
//
