package handlers

import (
	"encoding/json"
	"github.com/uvalib/entity-id-ws/entityidws/api"
	"log"
	"net/http"
	"strings"
)

func encodeStandardResponse(w http.ResponseWriter, status int) {

	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: http.StatusText(status)}); err != nil {
		log.Fatal(err)
	}
}

func encodeDetailsResponse(w http.ResponseWriter, status int, entity api.Request) {

	jsonAttributes(w)
	w.WriteHeader(status)
	if status == http.StatusOK {
		if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: http.StatusText(status), Details: &entity}); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: http.StatusText(status)}); err != nil {
			log.Fatal(err)
		}
	}
}

func encodeHealthCheckResponse(w http.ResponseWriter, status int, message string) {
	healthy := status == http.StatusOK
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.HealthCheckResponse{CheckType: api.HealthCheckResult{Healthy: healthy, Message: message}}); err != nil {
		log.Fatal(err)
	}
}

func encodeVersionResponse(w http.ResponseWriter, status int, version string) {
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.VersionResponse{Version: version}); err != nil {
		log.Fatal(err)
	}
}

func jsonAttributes(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func parameterOK(param string) bool {
	return len(strings.TrimSpace(param)) != 0
}
