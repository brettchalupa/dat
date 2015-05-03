package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "dat"
	app.Usage = "print the date and time"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{cli.Author{Name: "Brett Chalupa", Email: "brett@brettchalupa.com"}}
	app.Action = func(c *cli.Context) {
		const layout = "Jan 2, 2006 at 3:04pm (MST)"
		t := time.Now()
		fmt.Println(t.Format(layout))
	}
	app.Run(os.Args)
}
