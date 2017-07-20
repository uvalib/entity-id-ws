package authtoken

import (
	"entityidws/logger"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Validate(endpoint string, activity string, token string, timeout int) bool {

	url := fmt.Sprintf("%s/authorize/%s/%s/%s", endpoint, "entityidservice", activity, token)
	//log.Printf( "%s\n", url )

	start := time.Now()
	resp, _, errs := gorequest.New().
		SetDebug(false).
		Get(url).
		Timeout(time.Duration(timeout) * time.Second).
		End()
	duration := time.Since(start)

	if errs != nil {
		logger.Log(fmt.Sprintf("ERROR: token auth (%s) returns %s in %s", url, errs, duration))
		return false
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	logger.Log(fmt.Sprintf("Token auth (%s) returns http %d in %s", url, resp.StatusCode, duration))
	return resp.StatusCode == http.StatusOK
}
