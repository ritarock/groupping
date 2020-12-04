package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ritarock/groupping/lib/action"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "grouping",
		Usage: "Ping multiple targets",
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				action.Run(c.Args().Slice())
			} else {
				fmt.Println("The target is not set.")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
