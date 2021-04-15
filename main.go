package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
			&cli.StringFlag{
				Name:  "country",
				Value: "Malaysia",
				Usage: "country of origin",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Hamza"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "malay" {
				fmt.Println("Selamat Pagi ", name)
			} else {
				fmt.Println("Good morning ", name)
			}

			if c.String("country") == "malaysia" {
				fmt.Println("Breakfast today is nasi lemak kukus!")
			} else {
				fmt.Println("What do you like to eat?")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
