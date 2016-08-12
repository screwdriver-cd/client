package main

import(
	"os"

	"github.com/urfave/cli"
	"github.com/screwdriver-cd/client/commands"
)

func main(){
	app := cli.NewApp()
	app.Name = "Swagger"
	app.Usage = "Continuous Delivery"
	app.Commands = []cli.Command{
			{
					Name: "pipelines",
					Usage: "Options for pipelines",
					Subcommands: []cli.Command{
							{
									Name: "list",
									Usage: "list all pipelines",
									Action: command.PipelinesList,
							},
					},
			},
			{
					Name: "jobs",
					Usage: "List jobs",
					Subcommands: []cli.Command {
							{
									Name: "list",
									Usage: "List all jobs",
									Flags: []cli.Flag{
											cli.StringFlag{
												Name: "pipelineID, p",
												Usage: "Only show jobs for pipeline",
											},
									},
									Action: command.JobsList,
									ArgsUsage: "[num elements] [page number]",
							},
					},
			},
			{
					Name: "builds",
					Usage: "List Builds",
					Subcommands: []cli.Command{
						{
								Name: "list",
								Usage: "List all builds",
								Flags: []cli.Flag{
									cli.StringFlag{
										Name: "jobID, j",
										Usage: "Filter builds by job",
									},
									cli.StringFlag{
										Name: "status, s",
										Usage: "Only show builds of a specific status",
									},
								},
								Action: command.BuildList,
								ArgsUsage: "[num elements] [page number]",
						},
						{
								Name: "get",
								Usage: "Get a specific build",
								ArgsUsage: "<id>",
								Action: command.BuildsGetID,
						},
					},
			},
	}
	app.Run(os.Args)
}
