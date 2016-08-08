package command

import(
	"strconv"
	"fmt"
	"encoding/json"
	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
)

func formattedPrint(a ...interface{}){
	m, err := json.MarshalIndent(a, " ", "  ")
	if err != nil {
		fmt.Println(err)	
	}
	fmt.Println(string(m))
}

// PipelinesList handles the get endpoints for pipeliens
func PipelinesList(c *cli.Context) error {
	if len(c.Args()) == 0 {
		resp, err := sd.Default.V3.GetV3Pipelines(nil)
		if err != nil {
			return err	
		}
		formattedPrint(resp)
	} else if len(c.Args()) == 2 {
		args := c.Args()
		count, err := strconv.Atoi(args[0])
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		page, err := strconv.Atoi(args[1])
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		co := int64(count)
		p := int64(page)
		params := v3.NewGetV3PipelinesParams().WithCount(&co).WithPage(&p)
		resp, err := sd.Default.V3.GetV3Pipelines(params)
		formattedPrint(resp)
	} else {
		return cli.ShowSubcommandHelp(c)	
	}
	return nil
}
