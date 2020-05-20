package scraper

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

const (
	protocol = "https"
	baseUrl  = "race.netkeiba.com/race/shutuba.html?race_id="
)

// Get URL of horse data table
func BuildURL(raceID string) (url string) {
	url = fmt.Sprintf("%s://%s", protocol, baseUrl)
	url += raceID

	return url
}

func getRespBody(req *http.Request) (string, error) {
	httpCli := new(http.Client)
	resp, err := httpCli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	utfBody := transform.NewReader(bufio.NewReader(resp.Body), japanese.EUCJP.NewDecoder())
	body, err := ioutil.ReadAll(utfBody)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(body)

	return buf.String(), err
}
