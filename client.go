package main

import (
	"os"
	"strings"

	sd "github.com/screwdriver-cd/client/client"
	"github.com/screwdriver-cd/client/commands"
	filter "github.com/screwdriver-cd/client/commands/filters"
	"github.com/urfave/cli"
)

// CreateApp creates an instance of the Screwdriver App
func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Screwdriver Client"
	app.Usage = "Continuous Delivery With Screwdriver"
	app.Commands = []cli.Command{
		{
			Name:  "pipelines",
			Usage: "Screwdriver Pipelines",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List all pipelines",
					Action: func(c command.Context) error {
						resp, err := command.PipelinesList(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c.(*cli.Context))
						}
						command.FormattedPrint(resp)
						return nil
					},
					ArgsUsage: "[pagination count] [pagination page]",
				},
			},
		},
		{
			Name:  "jobs",
			Usage: "Screwdriver Jobs",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List all jobs",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "pipelineID, p",
							Usage: "Only show jobs for specified pipeline",
						},
					},
					Action: func(c command.Context) error {
						resp, err := command.JobsList(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c.(*cli.Context))
						}
						pipelineIDFlag := c.String("pipelineID")
						if strings.Compare(pipelineIDFlag, "") != 0 {
							resp = filter.JobsFilterPipeline(resp, pipelineIDFlag)
						}
						command.FormattedPrint(resp)
						return nil
					},
					ArgsUsage: "[pagination count] [pagination page]",
				},
			},
		},
		{
			Name:  "builds",
			Usage: "Screwdriver Builds",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List all builds",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "jobID, j",
							Usage: "Filter builds by jobID",
						},
						cli.StringFlag{
							Name:  "status, s",
							Usage: "Only show builds of a certain status",
						},
					},
					Action: func(c command.Context) error {
						resp, err := command.BuildsList(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c.(*cli.Context))
						}
						jobIDFlag := c.String("jobID")
						if strings.Compare(jobIDFlag, "") != 0 {
							resp = filter.BuildsFilterJobs(resp, jobIDFlag)
						}
						buildStatusFlag := c.String("status")
						if strings.Compare(buildStatusFlag, "") != 0 {
							resp = filter.BuildsFilterStatus(resp, buildStatusFlag)
						}
						command.FormattedPrint(resp)
						return nil
					},
					ArgsUsage: "[pagination count] [pagination page]",
				},
				{
					Name:      "get",
					Usage:     "Get a specific build",
					ArgsUsage: "<id>",
					Action: func(c command.Context) error {
						resp, err := command.BuildsGetID(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c.(*cli.Context))
						}
						command.FormattedPrint(resp)
						return nil
					},
				},
			},
		},
	}
	return app
}

func main() {
	app := CreateApp()
	app.Run(os.Args)
}
