package main

import(
	"os"

	"github.com/urfave/cli"
	"github.com/screwdriver-cd/client/commands"
)

func main(){
	app := cli.NewApp()
	app.Name = "Swagger"
	app.Usage = "wat"
	idFlag := []cli.Flag{
			cli.StringFlag{
				Name: "pipelineID",
				Value: "",
				Usage: "Filter for pipeline",
			},	
	}
	buildFlags := []cli.Flag{
			cli.StringFlag{
				Name: "jobID",
				Value: "",
				Usage: "Filter by Job ID",
			},
			cli.StringFlag{
				Name: "status",
				Value: "",
				Usage: "Filter by status",
			},
	}
	app.Commands = []cli.Command{
			{
					Name: "pipelines-list",
					Usage: "List all pipelines",
					Action: command.PipelinesList,
			},	
			{
					Name: "jobs-list",
					Usage: "List jobs",
					Action: command.JobsList,
					Flags: idFlag,
			},
			{
					Name: "builds-list",
					Usage: "List Builds",
					Action: command.BuildList,
					Flags: buildFlags,
			},
	}
	app.Run(os.Args)
}
