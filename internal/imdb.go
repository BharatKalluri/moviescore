package getratings

import (
	"strings"
)

//ImdbMovie structure for parsing the json from omdbapi
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

//GetImdbRatings function takes the moviename as argument and returns the json result from the omdbapi
func GetImdbRatings(mname string) ImdbMovie {
	movieName := strings.Replace(mname, " ", "+", 9)
	movieInfo := new(ImdbMovie)
	err := GetJSON("http://www.omdbapi.com/?t="+movieName+"&plot=full", movieInfo)
	LogError(err)
	return *movieInfo
}
