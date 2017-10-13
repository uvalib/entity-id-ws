package ezid

import (
   "entityidws/api"
   "entityidws/config"
   "entityidws/logger"
   "fmt"
   "github.com/parnurzeal/gorequest"
   "io"
   "io/ioutil"
   "net/http"
   "time"
)

// status for the EZID objects

// StatusPublic -- the item is public
const StatusPublic = "public"
// StatusReserved -- the item is reserved
const StatusReserved = "reserved"
// StatusUnavailable -- the item is unavailable
const StatusUnavailable = "unavailable|withdrawn by Library"

//
// GetDoi -- get entity details when provided a DOI
//
func GetDoi(doi string) (api.Request, int) {

   // construct target URL
   url := fmt.Sprintf("%s/id/%s", config.Configuration.EzidServiceURL, doi)

   // issue the request
   start := time.Now()
   resp, responseBody, errs := gorequest.New().
      SetDebug(config.Configuration.Debug).
      Get(url).
      Timeout(time.Duration(config.Configuration.Timeout) * time.Second).
      End()
   duration := time.Since(start)

   // check for errors
   if errs != nil {
      logger.Log(fmt.Sprintf("ERROR: service (%s) returns %s in %s", url, errs, duration))
      return blankResponse(), http.StatusInternalServerError
   }

   defer io.Copy(ioutil.Discard, resp.Body)
   defer resp.Body.Close()

   logger.Log(fmt.Sprintf("Service (%s) returns http %d in %s", url, resp.StatusCode, duration))

   // check the response body for errors
   if !statusIsOk(responseBody) {
      logger.Log(fmt.Sprintf("Error response body: [%s]", responseBody))
      return blankResponse(), http.StatusBadRequest
   }

   // all good...
   return makeEntityFromBody(responseBody), http.StatusOK
}

//
// CreateDoi -- Create a new entity; we may or may not have complete entity details
//
func CreateDoi(shoulder string, request api.Request, status string) (api.Request, int) {

   // log if necessary
   logRequest(request)

   // construct target URL
   url := fmt.Sprintf("%s/shoulder/%s", config.Configuration.EzidServiceURL, shoulder)

   // build the request body
   requestBody, err := makeBodyFromRequest(request, status)

   // check for errors
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: creating service payload %s", err))
      return blankResponse(), http.StatusBadRequest
   }

   // issue the request
   start := time.Now()
   resp, responseBody, errs := gorequest.New().
      SetDebug(config.Configuration.Debug).
      SetBasicAuth(config.Configuration.EzidUser, config.Configuration.EzidPassphrase).
      Post(url).
      Send(requestBody).
      Timeout(time.Duration(config.Configuration.Timeout)*time.Second).
      Set("Content-Type", "text/plain").
      End()
   duration := time.Since(start)

   // check for errors
   if errs != nil {
      logger.Log(fmt.Sprintf("ERROR: service (%s) returns %s in %s", url, errs, duration))
      return blankResponse(), http.StatusInternalServerError
   }

   defer io.Copy(ioutil.Discard, resp.Body)
   defer resp.Body.Close()

   logger.Log(fmt.Sprintf("Service (%s) returns http %d in %s", url, resp.StatusCode, duration))

   // check the response body for errors
   if !statusIsOk(responseBody) {
      logger.Log(fmt.Sprintf("Error response body: [%s]", responseBody))
      logger.Log(fmt.Sprintf("Original request body: [%s]", requestBody))
      return blankResponse(), http.StatusBadRequest
   }

   // all good...
   return makeEntityFromBody(responseBody), http.StatusOK
}

//
// UpdateDoi -- Update an existing DOI to match the provided entity
//
func UpdateDoi(request api.Request, status string) int {

   // log if necessary
   logRequest(request)

   // construct target URL
   url := fmt.Sprintf("%s/id/%s", config.Configuration.EzidServiceURL, request.ID)

   // build the request body
   requestBody, err := makeBodyFromRequest(request, status)

   // check for errors
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: creating service payload %s", err))
      return http.StatusBadRequest
   }

   // issue the request
   start := time.Now()
   resp, responseBody, errs := gorequest.New().
      SetDebug(config.Configuration.Debug).
      SetBasicAuth(config.Configuration.EzidUser, config.Configuration.EzidPassphrase).
      Post(url).
      Send(requestBody).
      Timeout(time.Duration(config.Configuration.Timeout)*time.Second).
      Set("Content-Type", "text/plain").
      End()
   duration := time.Since(start)

   // check for errors
   if errs != nil {
      logger.Log(fmt.Sprintf("ERROR: service (%s) returns %s in %s", url, errs, duration))
      return http.StatusInternalServerError
   }

   defer io.Copy(ioutil.Discard, resp.Body)
   defer resp.Body.Close()

   logger.Log(fmt.Sprintf("Service (%s) returns http %d in %s", url, resp.StatusCode, duration))

   // check the response body for errors
   if !statusIsOk(responseBody) {
      logger.Log(fmt.Sprintf("Error response body: [%s]", responseBody))
      logger.Log(fmt.Sprintf("Original request body: [%s]", requestBody))
      return http.StatusBadRequest
   }

   // all good...
   return http.StatusOK
}

//
// DeleteDoi -- delete entity details when provided a DOI
//
func DeleteDoi(doi string) int {

   // construct target URL
   url := fmt.Sprintf("%s/id/%s", config.Configuration.EzidServiceURL, doi)

   // issue the request
   start := time.Now()
   resp, responseBody, errs := gorequest.New().
      SetDebug(config.Configuration.Debug).
      SetBasicAuth(config.Configuration.EzidUser, config.Configuration.EzidPassphrase).
      Delete(url).
      Timeout(time.Duration(config.Configuration.Timeout) * time.Second).
      End()
   duration := time.Since(start)

   // check for errors
   if errs != nil {
      logger.Log(fmt.Sprintf("ERROR: service (%s) returns %s in %s", url, errs, duration))
      return http.StatusInternalServerError
   }

   defer io.Copy(ioutil.Discard, resp.Body)
   defer resp.Body.Close()

   logger.Log(fmt.Sprintf("Service (%s) returns http %d in %s", url, resp.StatusCode, duration))

   // check the response body for errors
   if !statusIsOk(responseBody) {
      logger.Log(fmt.Sprintf("Error response body: [%s]", responseBody))
      return http.StatusBadRequest
   }

   // all good...
   return http.StatusOK
}

//
// GetStatus -- get the status of the endpoint
//
func GetStatus() int {

   // construct target URL
   url := fmt.Sprintf("%s/status", config.Configuration.EzidServiceURL)

   // issue the request
   start := time.Now()
   resp, responseBody, errs := gorequest.New().
      SetDebug(config.Configuration.Debug).
      Get(url).
      Timeout(time.Duration(config.Configuration.Timeout) * time.Second).
      End()
   duration := time.Since(start)

   // check for errors
   if errs != nil {
      logger.Log(fmt.Sprintf("ERROR: service (%s) returns %s in %s", url, errs, duration))
      return http.StatusInternalServerError
   }

   defer io.Copy(ioutil.Discard, resp.Body)
   defer resp.Body.Close()

   logger.Log(fmt.Sprintf("Service (%s) returns http %d in %s", url, resp.StatusCode, duration))

   // check the response body for errors
   if !statusIsOk(responseBody) {
      logger.Log(fmt.Sprintf("Error response body: [%s]", responseBody))
      return http.StatusBadRequest
   }

   // all good...
   return http.StatusOK
}

//
// end of file
//
