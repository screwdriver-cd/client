package main

import(
	"os"

	"github.com/urfave/cli"
	"github.com/screwdriver-cd/client/commands"
	filter "github.com/screwdriver-cd/client/commands/filters"
)

func StartApp() *cli.App { 
	app := cli.NewApp()	
	app.Name = "Screwdriver Client"
	app.Usage = "Continuous Delivery With Screwdriver"
	app.Commands = []cli.Command{
		{
			Name: "pipelines",
			Usage: "Screwdriver Pipelines",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Usage: "List all pipelines",
					Action: func(c *cli.Context) error {
						resp, err := command.PipelinesList(c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)	
						}
						command.FormattedPrint(resp)
						return nil
					},
				},
			},	
		},
		{
			Name: "jobs",
			Usage: "Screwdriver Jobs",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Usage: "List all jobs",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "pipelineID, p",
							Usage: "Only show jobs for specified pipeline",
						},
					},
					Action: func(c *cli.Context) error {
						resp, err := command.JobsList(c)
						if err != nil {
							return cli.ShowSubcommandHelp(c)	
						}
						resp = filter.JobsFilterPipeline(resp, c)
						command.FormattedPrint(resp)
						return nil
					},
				ArgsUsage: "[num elements] [page number]",
				},
			},
		},
		{
			Name: "builds",
			Usage: "Screwdriver Builds",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Usage: "List all builds",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "jobID, j",
							Usage: "Filter builds by jobID",
						},
						cli.StringFlag{
							Name: "status, s",
							Usage: "Only show builds of a certain status",
						},
					},
					Action: func(c *cli.Context) error{
						resp, err := command.BuildsList(c)	
						if err != nil {
							return cli.ShowSubcommandHelp(c)	
						}
						resp = filter.BuildsFilterJobs(resp,c)
						resp = filter.BuildsFilterStatus(resp,c)
						command.FormattedPrint(resp)
						return nil
					},
					ArgsUsage: "[num elements] [page number]",
				},
				{
						Name: "get",
						Usage: "Get a specific build",
						ArgsUsage: "<id>",
						Action: func(c *cli.Context) error {
							resp, err := command.BuildsGetID(c)
							if err != nil {
								return cli.ShowSubcommandHelp(c)
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

func main(){
	app := StartApp()	
	app.Run(os.Args)
}
