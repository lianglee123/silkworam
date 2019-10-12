package main


import (
	"github.com/urfave/cli"
	"os"
	"log"
)

var (
	cliApp *cli.App
)

func init() {
	// Initialise a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "demo"
	cliApp.Usage = "demo usage"
	cliApp.Author = "joyLee"
	cliApp.Email = "1025012627@qq.com"
	cliApp.Version = "0.0.1"
}

func main() {
	cliApp.Commands = []cli.Command{
		{
			Name: "run",
			Usage: "run http server and grpc server",
			Action: func(c *cli.Context) {

			},
		},

		{
			Name: "config",
			Usage: "show all config",
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}