package main

import (
	"fmt"
	"github.com/bharatkalluri/moviescore/getratings"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "MovieScore"
	app.Usage = "A cli utility for showing Movie Ratings!"
	app.UsageText = "MovieScore <Movie name here> (Please have quotes on either side if the movie name has spaces)"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "year, y",
			Usage: "year of release, useful for rt ratings",
		},
		cli.StringFlag{
			Name:  "reviews, r",
			Usage: "Reviews, will be pulled from Rotten Tomatoes",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("reviews") != "" {
			AsciiPoster()
			fmt.Println(chalk.Green, "Reviews from rt", chalk.Reset)
			fmt.Println(chalk.Magenta, "-------------------", chalk.Reset)
			getratings.RtReviewScraper(c.String("reviews"), c.String("year"))
			return nil
		} else if len(c.Args()) > 1 {
			fmt.Println("See the help by typing 'moviescore -h'")
		} else if len(c.String("year")) == 4 {
			PrettyPrinter(c.Args().Get(0), c.String("year"))
		} else {
			PrettyPrinter(c.Args().Get(0), "")
		}
		fmt.Println(c.String("reviews"))
		getratings.RtReviewScraper(c.String("reviews"), "")
		return nil
	}
	app.Run(os.Args)
}

func PrettyPrinter(MovieName string, year string) {
	RtRating := getratings.RtScraper(MovieName, year)
	ImdbRatings := getratings.GetImdbRatings(MovieName)
	IntRtRatings, err := strconv.Atoi(RtRating)
	AsciiPoster()
	if IntRtRatings == -1 && len(ImdbRatings.Title) == 0 {
		fmt.Println("The Movie Does not seem to exist!")
		fmt.Println("Tip: If you are using spaces in your film name, enclose the movie name in double quotes!")
	} else {
		fmt.Println(chalk.Magenta, "Movie Name: "+ImdbRatings.Title)
		fmt.Println(chalk.Magenta, "Director: "+ImdbRatings.Director)
		fmt.Println(chalk.Magenta, "Cast: "+ImdbRatings.Actors)
		fmt.Println(chalk.Magenta, "Year: "+ImdbRatings.Year)
		fmt.Println(chalk.Magenta, "Released: "+ImdbRatings.Released)
		fmt.Println(chalk.Magenta, "Rated: "+ImdbRatings.Rated)
		fmt.Println(chalk.Magenta, "Genre: "+ImdbRatings.Genre)
		fmt.Println(chalk.Magenta, "Poster: "+ImdbRatings.Poster)
		fmt.Println(chalk.Magenta, "Metascore Rated: "+ImdbRatings.Metascore)
		fmt.Println(chalk.Magenta, "Awards: "+ImdbRatings.Awards)
		fmt.Println(chalk.Magenta, "Plot: "+ImdbRatings.Plot)
		fmt.Println(" Ratings from IMDB and Rotten Tomatoes-")
		fmt.Println(chalk.Magenta, chalk.Underline.TextStyle("IMDB Rating: "+ImdbRatings.ImdbRating))
		if IntRtRatings == -1 && err == nil {
			fmt.Println(chalk.Red, "There seems to be a problem with rt, try with the year argument!")
		} else if IntRtRatings > 60 && err == nil {
			fmt.Println(chalk.Red, chalk.Underline.TextStyle("Rotten Tomatoes Rating: "+RtRating+"% (Certified Fresh!)"))
		} else {
			fmt.Println(chalk.Green, chalk.Underline.TextStyle("Rotten Tomatoes Rating: "+RtRating+"% (Rotten!)"))
		}
	}
}

func AsciiPoster() {
	fmt.Println(chalk.Cyan, `
------------------------------------------------------
  __  __            _         _____                    
 |  \/  |          (_)       / ____|                   
 | \  / | _____   ___  ___  | (___   ___ ___  _ __ ___ 
 | |\/| |/ _ \ \ / / |/ _ \  \___ \ / __/ _ \| '__/ _ \
 | |  | | (_) \ V /| |  __/  ____) | (_| (_) | | |  __/
 |_|  |_|\___/ \_/ |_|\___| |_____/ \___\___/|_|  \___|
------------------------------------------------------
	`)
}
