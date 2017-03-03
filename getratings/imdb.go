package getratings

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ImdbMovie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Genre      string
	Director   string
	Actors     string
	Plot       string
	Awards     string
	Poster     string
	Metascore  string
	ImdbRating string `json:"imdbRating"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func GetImdbRatings(mname string) ImdbMovie {
	movieName := strings.Replace(mname, " ", "+", 9)
	movieInfo := new(ImdbMovie)
	getJSON("http://www.omdbapi.com/?t="+movieName+"&plot=full", movieInfo)
	return *movieInfo
}
