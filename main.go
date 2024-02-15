package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/urfave/cli/v2"
)

func parseFeed(targetFeed string) {
	feed, err := gofeed.NewParser().ParseURL(targetFeed)
	if err != nil {
		panic(err)
	}

	fmt.Println(feed.Title)
	for _, item := range feed.Items {
		fmt.Printf("\t%s -> %s\n", item.Title, item.Link)
	}
}

func main() {
	var targetFeed string
	app := &cli.App{
		Name:  "Feed Parser",
		Usage: "Parse and display RSS feeds",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "feed",
				Value:       "https://feeds.bbci.co.uk/news/rss.xml",
				Usage:       "Feed to parse",
				Destination: &targetFeed,
			},
		},
		Action: func(*cli.Context) error {
			parseFeed(targetFeed)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
