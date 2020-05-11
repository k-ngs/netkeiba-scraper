package scraper

import (
	"fmt"
	"net/http"
)

const (
	protocol = "https"
	baseUrl  = "race.netkeiba.com/race/shutuba.html?race_id="
)

// Get URL of horse data table
func BuildURL(year, course, count, days, raceNum string) (url string) {
	url = fmt.Sprintf("%s://%s", protocol, baseUrl)
	raceId := year + course + count + days + raceNum
	url += raceId

	return url
}

func sendReq(req *http.Request) (resp *http.Response, err error) {
	httpCli := new(http.Client)
	resp, err = httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	return
}
