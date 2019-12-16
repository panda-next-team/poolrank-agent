package main

import (
	"fmt"
	gofixerio "github.com/LordotU/go-fixerio"
	fixer_io "github.com/panda-next-team/poolrank-agent/internal/app/job/fixer.io"
	"github.com/panda-next-team/poolrank-agent/internal/pkg"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	var config fixer_io.FetchRateLatestConfig
	app := cli.NewApp()
	app.Name = "poolrank-agent-fixer-fetch-rate-latest"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2019 poolrank"
	app.Usage = "fetch rate latest data from fixer.io"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Alex Zhou",
			Email: "333266664@qq.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dbHost, dbh",
			Value:       "127.0.0.1",
			Usage:       "mysql host",
			EnvVar:      "AGENT_DB_HOST",
			Destination: &config.Mysql.Host,
		},
		cli.StringFlag{
			Name:        "dbUser, dbu",
			Value:       "root",
			Usage:       "mysql user",
			EnvVar:      "AGENT_DB_USER",
			Destination: &config.Mysql.User,
		},
		cli.StringFlag{
			Name:        "dbPassword, dbp",
			Value:       "",
			Usage:       "mysql password",
			EnvVar:      "AGENT_DB_PASSWORD",
			Destination: &config.Mysql.Password,
		},
		cli.StringFlag{
			Name:        "dbName, dbn",
			Value:       "pr_agent",
			Usage:       "mysql database name",
			EnvVar:      "AGENT_DB_NAME",
			Destination: &config.Mysql.Database,
		},
		cli.IntFlag{
			Name:        "dbPort",
			Value:       3306,
			Usage:       "mysql db port",
			EnvVar:      "AGENT_DB_PORT",
			Destination: &config.Mysql.Port,
		},
		cli.StringFlag{
			Name:        "dbPrefix",
			Value:       "pa_",
			Usage:       "mysql db prefix",
			EnvVar:      "AGENT_DB_PREFIX",
			Destination: &config.Mysql.Prefix,
		},
		cli.StringFlag{
			Name:        "apiKey, key",
			Value:       "",
			Usage:       "fixer.io api key",
			EnvVar:      "AGENT_FIXER_KEY",
			Destination: &config.Fixer.APIKey,
		},
		cli.StringSliceFlag{
			Name: "symbols",
			Usage: "currency symbols",
			EnvVar: "AGENT_CURRENCY_SYMBOLS",
			Value:  &cli.StringSlice{"USD", "CNY"},
		},
	}

	app.Before = func(c *cli.Context) error {
		config.Symbols = c.StringSlice("symbols")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run grpc server",
			Action: func(c *cli.Context) {
				engine, engineErr := pkg.NewMysqlEngine(pkg.DriverName, pkg.GenerateMysqlSource(
					config.Mysql.User, config.Mysql.Password, config.Mysql.Host,
					config.Mysql.Port, config.Mysql.Database), config.Mysql.Prefix)

				if engineErr != nil {
					log.Fatal(fmt.Sprintf("db init err %s", engineErr))
				}

				fixerIO, fixerErr := gofixerio.New(config.Fixer.APIKey, "EUR", false)
				if fixerErr != nil {
					log.Fatal(fmt.Sprintf("fixer client init err %s", fixerErr))
				}

				job := &fixer_io.FetchRateLatestJob{engine, fixerIO}
				if err := job.Run(config.Symbols); err != nil {
					log.Fatal(err)
				}

				log.Println("run success")
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
