package scraper

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func newGetExampleResp() (resp *http.Response) {
	resp = &http.Response{
		Body: NewBodyWithString(`<?xml version="1.0" encoding="UTF-8"?>
<example></example>
 `),
	}
	return
}

func TestBuildURL(t *testing.T) {
	var (
		year    = "2020"
		course  = "05"
		count   = "03"
		days    = "01"
		raceNum = "08"
	)
	expect := "https://race.netkeiba.com/race/shutuba.html?race_id=202005030108"
	result := BuildURL(year, course, count, days, raceNum)

	assertEquals(t, result, expect)
}

func TestSendReq(t *testing.T) {
	var (
		year    = "2020"
		course  = "05"
		count   = "03"
		days    = "01"
		raceNum = "08"
	)
	expect := newGetExampleResp()
	target := BuildURL(year, course, count, days, raceNum)
	req, err := http.NewRequest("GET", target, nil)
	result, err := sendReq(req)
	if err != nil {
		t.Errorf("HTTP request is failed\n")
	}
	assertEquals(t, result, expect)
}

func assertEquals(t *testing.T, result, expect interface{}) {
	if result != expect {
		t.Errorf("result: %v, expect: %v\n", result, expect)
	}
}

type DummyResponseBody struct {
	data io.Reader
}

func NewBodyWithString(data string) *DummyResponseBody {
	return &DummyResponseBody{data: strings.NewReader(data)}
}
