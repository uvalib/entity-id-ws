package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/uvalib/entity-id-ws/entityidws/api"
	"github.com/uvalib/entity-id-ws/entityidws/authtoken"
	"github.com/uvalib/entity-id-ws/entityidws/config"
	"github.com/uvalib/entity-id-ws/entityidws/idservice"
	"github.com/uvalib/entity-id-ws/entityidws/logger"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//
// IDUpdate -- the update by ID request handler
//
func IDUpdate(w http.ResponseWriter, r *http.Request) {

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

	decoder := json.NewDecoder(r.Body)
	request := api.Request{}

	if err := decoder.Decode(&request); err != nil {
		logger.Log(fmt.Sprintf("ERROR: decoding request payload %s", err))
		encodeStandardResponse(w, http.StatusBadRequest)
		return
	}

	defer io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	request.ID = doi
	status := idservice.UpdateDoi(request, idservice.StatusPublic)
	encodeStandardResponse(w, status)
}

//
// end of file
//
