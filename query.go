package prometheus_query

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const sevenDaysSeconds = 86400 * 7

func Query(URL string, token string, queries []string, plotterChan chan<- http.Response) {
	startTime := time.Now().Unix() - sevenDaysSeconds
	endTime := time.Now().Unix()
	startTimeString := strconv.FormatInt(startTime, 10)
	endTimeString := strconv.FormatInt(endTime, 10)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	client := &http.Client{
		Transport: transCfg,
	}

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		log.Fatalf("Error building Request object: %v", err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("query", "up")
	q.Add("start", startTimeString)
	q.Add("end", endTimeString)
	req.URL.RawQuery = q.Encode()

	log.Print(req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error while performing request: %v", err)
		os.Exit(1)
	}

	plotterChan <- *resp
}

func parseResponse(response *http.Response) string {
	locationHeader := response.Header.Get("metric")
	accessToken_raw := strings.Split(locationHeader, "access_token=")
	accessToken := strings.Split(accessToken_raw[len(accessToken_raw)-1], "&")[0]
	return accessToken
}
