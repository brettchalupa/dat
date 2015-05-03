package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "tad"
	app.Usage = "print the date and time"
	app.Action = func(c *cli.Context) {
		const layout = "Jan 2, 2006 at 3:04pm (MST)"
		t := time.Now()
		fmt.Println(t.Format(layout))
	}
	app.Run(os.Args)
}
