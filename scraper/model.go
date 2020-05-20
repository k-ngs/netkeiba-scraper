package scraper

type raceData struct {
	id       string
	raceName string
	horses   []horse
}

type horse struct {
	name    string
	waku    string
	num     string
	age     string
	jockey  string
	stable  string
	trainer string
}
