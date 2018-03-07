package handlers

import (
	"encoding/json"
	"entityidws/api"
	"entityidws/authtoken"
	"entityidws/config"
	"entityidws/idservice"
	"entityidws/logger"
	"fmt"
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
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, "update", token, config.Configuration.ServiceTimeout) == false {
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
