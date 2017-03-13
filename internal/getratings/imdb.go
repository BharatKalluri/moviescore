package getratings

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ttacon/chalk"
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
	ImdbID     string
}

//GetImdbRatings function takes the moviename as argument and returns the json result from the omdbapi
func GetImdbRatings(mname string) ImdbMovie {
	movieName := strings.Replace(mname, " ", "+", 9)
	movieInfo := new(ImdbMovie)
	err := GetJSON("http://www.omdbapi.com/?t="+movieName+"&plot=full", movieInfo)
	LogError(err)
	return *movieInfo
}

// Scrapes IMDB Parental Ratings
func GetImdbParentsGuide(MovieName string) {
	fmt.Println(chalk.Magenta, "Parental Guide:")
	movieName := strings.Replace(MovieName, " ", "+", 9)
	movieInfo := new(ImdbMovie)
	err := GetJSON("http://www.omdbapi.com/?t="+movieName+"&plot=full", movieInfo)
	LogError(err)
	ImdbID := movieInfo.ImdbID
	urlis := "http://www.imdb.com/title/" + ImdbID + "/parentalguide"
	fmt.Println("Reference Url: ", urlis)
	doc, err := goquery.NewDocument(urlis)
	if err != nil {
		fmt.Println(err)
	}
	SexAndNudity := doc.Find("#swiki\\.2\\.1\\.1").Text()
	ViolenceAndGore := doc.Find("#swiki\\.2\\.2\\.1").Text()
	Profanity := doc.Find("#swiki\\.2\\.3\\.1").Text()
	Alcohol := doc.Find("#swiki\\.2\\.4\\.1").Text()
	Intense := doc.Find("#swiki\\.2\\.5\\.1").Text()
	SexAndNudity = DataCleaner(SexAndNudity)
	ViolenceAndGore = DataCleaner(ViolenceAndGore)
	Profanity = DataCleaner(Profanity)
	Alcohol = DataCleaner(Alcohol)
	Intense = DataCleaner(Intense)
	fmt.Println(chalk.Green.NewStyle().WithTextStyle(chalk.Bold), "Sex and Nudity", chalk.Reset)
	fmt.Println(SexAndNudity)
	fmt.Println(chalk.Green.NewStyle().WithTextStyle(chalk.Bold), "Violence And Gore", chalk.Reset)
	fmt.Println(ViolenceAndGore)
	fmt.Println(chalk.Green.NewStyle().WithTextStyle(chalk.Bold), "Profanity", chalk.Reset)
	fmt.Println(Profanity)
	fmt.Println(chalk.Green.NewStyle().WithTextStyle(chalk.Bold), "Alcohol/Drugs/Smoking", chalk.Reset)
	fmt.Println(Alcohol)
	fmt.Println(chalk.Green.NewStyle().WithTextStyle(chalk.Bold), "Frightening/Intense Scenes", chalk.Reset)
	fmt.Println(Intense)
}

func DataCleaner(data string) string {
	data = strings.Replace(data, "<br> \n", "", -1)
	data = strings.TrimSpace(data)
	return data
}
