package prometheus_query

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"../util"
)

const sevenDaysSeconds = 86400 * 7

func Query(Url string, token string, queries []util.PromQuery, plotterChan chan<- util.QueryResult) {
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

		setQueryStartEnd(req, &query, startTimeString, endTimeString)

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

		result := util.QueryResult{
			PrettyName: query.PrettyName,
			Result:     string(bodyBytes),
		}

		plotterChan <- result
	}
}

func setQueryStartEnd(httpRequest *http.Request, promQuery *util.PromQuery, defaultStart string, defaultEnd string) {
	q := httpRequest.URL.Query()

	q.Add("query", promQuery.Expression)

	if promQuery.Start == "" {
		q.Add("start", defaultStart)
	} else {
		q.Add("start", promQuery.Start)
	}

	if promQuery.End == "" {
		q.Add("end", defaultEnd)
	} else {
		q.Add("end", promQuery.End)
	}

	httpRequest.URL.RawQuery = q.Encode()
}
