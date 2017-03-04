package main

import (
	"fmt"
	"github.com/bharatkalluri/moviescore/getratings"
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
			getratings.AsciiPoster()
			fmt.Println(chalk.Green, "Reviews from rt", chalk.Reset)
			fmt.Println(chalk.Magenta, "-------------------", chalk.Reset)
			getratings.RtReviewScraper(c.String("reviews"), c.String("year"))
			return nil
		} else if len(c.Args()) > 1 || len(c.Args()) == 0 {
			fmt.Println("See the help by typing 'moviescore -h'")
			return nil
		} else if len(c.String("year")) == 4 {
			getratings.PrettyPrinter(c.Args().Get(0), c.String("year"))
		} else {
			getratings.PrettyPrinter(c.Args().Get(0), "")
		}
		fmt.Println(c.String("reviews"))
		return nil
	}
	app.Run(os.Args)
}
