package handlers

import (
	"encoding/json"
	"github.com/uvalib/entity-id-ws/entityidws/api"
	"github.com/uvalib/entity-id-ws/entityidws/authtoken"
	"github.com/uvalib/entity-id-ws/entityidws/config"
	"github.com/uvalib/entity-id-ws/entityidws/idservice"
	"github.com/uvalib/entity-id-ws/entityidws/logger"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//
// IDCreate -- the create by ID request handler
//
func IDCreate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	shoulder := vars["shoulder"]
	token := r.URL.Query().Get("auth")

	// validate inbound parameters
	if parameterOK(shoulder) == false || parameterOK(token) == false {
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, "create", token, config.Configuration.ServiceTimeout) == false {
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

	entity, status := idservice.CreateDoi(shoulder, request, idservice.StatusReserved)
	encodeDetailsResponse(w, status, entity)
}

//
// end of file
//
