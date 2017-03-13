package main

import (
	"fmt"
	"github.com/bharatkalluri/MovieScore/internal/getratings"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "MovieScore"
	app.Usage = "A cli utility for showing Movie Ratings!"
	app.UsageText = "MovieScore <Movie name here> (Please have quotes on either side if the movie name has spaces)"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "parentalguide, pg",
			Usage: "Parental Guide from IMDB, Does not need the year argument.",
		},
		cli.BoolFlag{
			Name:  "reviews, r",
			Usage: "Reviews, will be pulled from Rotten Tomatoes",
		},
		cli.StringFlag{
			Name:  "year, y",
			Usage: "year of release, useful for rt ratings",
		},
	}

	app.Action = func(c *cli.Context) error {

		getratings.ASCIIPoster()
		if c.Bool("parentalguide") == true && c.Bool("reviews") == true {
			fmt.Println(chalk.Red, "One option at a time, cant have -pg and -r flags at the same time!")
		}
		if c.Bool("parentalguide") == true {
			getratings.GetImdbParentsGuide(c.Args().Get(0))
		} else if c.Bool("reviews") == true {
			getratings.RtReviewScraper(c.Args().Get(0), c.String("year"))
		} else {
			getratings.PrettyPrinter(c.Args().Get(0), c.String("year"))
		}
		return nil
	}
	app.Run(os.Args)
}
