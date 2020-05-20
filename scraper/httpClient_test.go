package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBuildURL(t *testing.T) {
	var raceID = "202005020811"
	expect := "https://race.netkeiba.com/race/shutuba.html?race_id=202005020811"
	result := BuildURL(raceID)

	assertEquals(t, result, expect)
}

func TestGetRespBody(t *testing.T) {
	r := raceData{
		id: "202005020811",
	}
	var mockFileName = "./testFile/202005020811.html"

	expect, err := ioutil.ReadFile(mockFileName)
	target := BuildURL(r.id)
	req, err := http.NewRequest("GET", target, nil)
	result, err := getRespBody(req)
	if err != nil {
		t.Errorf("HTTP request is failed\n")
	}

	assertEquals(t, result, string(expect))
}

func TestGetRaceTable(t *testing.T) {
	expect := &raceData{
		id:       "202005020811",
		raceName: "ヴィクトリアM",
	}

	result, err := newRaceData(expect.id)
	if err != nil {
		t.Errorf("Failed to initiate race data\n")
	}
	fmt.Println(result)
	assertEquals(t, result, expect)
}

func assertEquals(t *testing.T, result, expect interface{}) {
	if result != expect {
		t.Errorf("result: %v, expect: %v\n", result, expect)
	}
}
