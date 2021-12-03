package prometheus_query

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"../util"
)

const sevenDaysSeconds = 86400 * 7

func Query(Url string, token string, queries []util.PromQuery, plotterChan chan<- string) {
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

	for _, query := range queries {

		req, err := http.NewRequest("GET", Url, nil)

		if err != nil {
			log.Fatalf("Error building Request object: %v", err)
			os.Exit(1)
		}

		q := req.URL.Query()
		if query.Start == "" {
			q.Add("start", startTimeString)
		} else {
			q.Add("start", query.Start)
		}

		q.Add("query", query.Expression)
		if query.End == "" {
			q.Add("end", endTimeString)
		} else {
			q.Add("end", query.End)
		}

		req.URL.RawQuery = q.Encode()

		log.Print(req.URL.String())

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error while performing request: %v", err)
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error while reading response: %v", err)
			os.Exit(1)
		}
		resp.Body.Close()
		bodyString := string(bodyBytes)

		plotterChan <- bodyString
	}
}

func parseResponse(response *http.Response) string {
	locationHeader := response.Header.Get("metric")
	accessToken_raw := strings.Split(locationHeader, "access_token=")
	accessToken := strings.Split(accessToken_raw[len(accessToken_raw)-1], "&")[0]
	return accessToken

}
