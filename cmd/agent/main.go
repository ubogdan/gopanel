package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"

	"github.com/ubogdan/gopanel/model"
)

const (
	daemonFlag = "d"
	debugFlag  = "debug"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-panel-agent"
	app.Version = model.Version().String()
	app.Action = startService

	daemonCliFlag := cli.BoolFlag{ //nolint:exhaustivestruct
		Name: daemonFlag,
	}

	debugCliFlag := cli.BoolFlag{ //nolint:exhaustivestruct
		Name: debugFlag,
	}

	app.Flags = []cli.Flag{
		&daemonCliFlag,
		&debugCliFlag,
	}

	app.Run(os.Args)
}

func startService(c *cli.Context) error {
	return nil
}
