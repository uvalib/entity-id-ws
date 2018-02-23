package client

import (
	"encoding/json"
	"entityidws/api"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var debugHTTP = false
var serviceTimeout = 5

//
// HealthCheck -- calls the service health check method
//
func HealthCheck(endpoint string) int {

	url := fmt.Sprintf("%s/healthcheck", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode
}

//
// VersionCheck -- calls the service version check method
//
func VersionCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/version", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.VersionResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, ""
	}

	return resp.StatusCode, r.Version
}

//
// MetricsCheck -- calls the service metrics method
//
func MetricsCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/metrics", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(false).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode, body
}

//
// Statistics -- call the statistics method on the service
//
func Statistics(endpoint string) (int, *api.Statistics) {

	url := fmt.Sprintf("%s/statistics", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.StatisticsResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, &r.Details
}

//
// Get -- call the get method on the service
//
func Get(endpoint string, doi string, token string) (int, *api.Request) {

	url := fmt.Sprintf("%s/%s?auth=%s", endpoint, doi, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.StandardResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, r.Details
}

//
// Create -- call the create method on the service
//
func Create(endpoint string, shoulder string, entity api.Request, token string) (int, *api.Request) {

	url := fmt.Sprintf("%s/%s?auth=%s", endpoint, shoulder, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Post(url).
		Send(entity).
		Timeout(time.Duration(serviceTimeout)*time.Second).
		Set("Content-Type", "application/json").
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.StandardResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, r.Details
}

//
// Update -- call the update method on the service
//
func Update(endpoint string, entity api.Request, token string) int {

	url := fmt.Sprintf("%s/%s?auth=%s", endpoint, entity.ID, token)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Put(url).
		Send(entity).
		Timeout(time.Duration(serviceTimeout)*time.Second).
		Set("Content-Type", "application/json").
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode
}

//
// Delete -- call the delete method on the service
//
func Delete(endpoint string, doi string, token string) int {

	url := fmt.Sprintf("%s/%s?auth=%s", endpoint, doi, token)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Delete(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode
}

//
// Revoke -- revoke an entity when provided a DOI
//
func Revoke(endpoint string, doi string, token string) int {

	// construct target URL
	url := fmt.Sprintf("%s/revoke/%s?auth=%s", endpoint, doi, token)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Put(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode
}

//
// end of file
//
