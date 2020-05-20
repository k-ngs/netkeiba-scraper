package scraper

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func newRaceData(raceID string) (*raceData, error) {
	d := &raceData{
		id: raceID,
	}
	err := getRaceTable(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func getRaceTable(d *raceData) error {
	var tmp horse
	target := BuildURL(d.id)
	req, _ := http.NewRequest("GET", target, nil)
	body, err := getRespBody(req)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(body))

	// Get race name
	raceName := doc.Find("div.RaceName").First()
	d.raceName = strings.TrimSpace(raceName.Text())

	// Get horse data
	horses := doc.Find("tr.HorseList")
	horses.Each(func(_ int, s *goquery.Selection) {
		tmp.name = strings.TrimSpace(s.Find("span.HorseName").Text())

		tmp.waku = s.Find("td").First().Text()
		numSelector := "td.Umaban" + tmp.waku + ".Txt_C"
		tmp.num = s.Find(numSelector).Text()

		tmp.age = s.Find("td.Barei").Text()
		tmp.jockey = strings.TrimSpace(s.Find("td.Jockey").Text())

		trainer := s.Find("td.Trainer")
		tmp.stable = trainer.Find("span").Text()
		tmp.trainer = trainer.Find("a").Text()

		d.horses = append(d.horses, tmp)

	})

	return nil
}
