package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "RabbitMQ Worker"
	app.Usage = "Register as a consumer for a given queue and proxy the messages onto a cli command."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config-file, c",
			Usage: "Relative or absolute path to config file `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("hello world")

		if err := InitConfig(c.String("config-file")); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return nil
	}

	app.Run(os.Args)
}
