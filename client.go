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
					Action: func(c *cli.Context) error {
						resp, err := command.PipelinesList(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
						}
						command.FormattedPrint(resp)
						return nil
					},
					ArgsUsage: "[pagination count] [pagination page]",
				},
				{
					Name:  "get",
					Usage: "Get a pipeline by ID",
					Action: func(c *cli.Context) error {
						resp, err := command.PipelinesGetID(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
						}
						command.FormattedPrint(resp)
						return nil
					},
					ArgsUsage: "<id>",
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
					Action: func(c *cli.Context) error {
						resp, err := command.JobsList(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
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
					Action: func(c *cli.Context) error {
						resp, err := command.BuildsList(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
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
					Action: func(c *cli.Context) error {
						resp, err := command.BuildsGetID(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
						}
						command.FormattedPrint(resp)
						return nil
					},
				},
				{
					Name:      "steps",
					Usage:     "Get a step record",
					ArgsUsage: "<id> <stepName>",
					Action: func(c *cli.Context) error {
						resp, err := command.BuildsGetStep(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
						}
						command.FormattedPrint(resp)
						return nil
					},
				},
				{
					Name:      "steps-log",
					Usage:     "Get the logs for a step",
					ArgsUsage: "<id> <stepName>",
					Action: func(c *cli.Context) error {
						resp, err := command.BuildsGetStepLogs(sd.Default, c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)
						}
						command.FormattedPrint(resp)
						return nil
					},
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "start, s",
							Usage: "Start the logs for the step from a specific number",
						},
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
