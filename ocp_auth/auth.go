package ocp_auth

import (
	"crypto/tls"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"
)

const OauthPath = "/oauth/authorize"
const response_type = "token"
const client_id = "openshift-challenging-client"

func GetOcpToken(ocpOauthUrl string, username string, password string) string {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	client := &http.Client{
		Transport: transCfg,
	}

	encodedAuth := basicAuth(username, password)
	req, err := http.NewRequest("GET", ocpOauthUrl+OauthPath, nil)

	if err != nil {
		log.Fatalf("Error building Request object: %v", err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("ResponseType", response_type)
	q.Add("ClientID", client_id)
	req.URL.RawQuery = q.Encode()

	log.Print(req.URL.String())

	req.Header.Add("Authorization", "Basic "+encodedAuth)
	req.Header.Add("X-CSRF-Token", "csrf-token")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error while performing request: %v", err)
		os.Exit(1)
	}

	return parseResponse(resp)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func parseResponse(response *http.Response) string {
	locationHeader := response.Header.Get("Location")
	accessToken_raw := strings.Split(locationHeader, "access_token=")
	accessToken := strings.Split(accessToken_raw[len(accessToken_raw)-1], "&")[0]
	return accessToken
}
