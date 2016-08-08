package command

import(
	"strconv"
	"strings"

	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)

// jobFilterPipeline filters jobs by pipelineID returns the original v3.GetV3JobsOK object
func jobsFilterPipeline(resp *v3.GetV3JobsOK, c *cli.Context) *v3.GetV3JobsOK {
	filters := c.String("pipelineID")
	if strings.Compare(filters, "") != 0 {
			res := model.ListOfJobs{}
			for _, element := range resp.Payload{
				if strings.Compare(filters, element.PipelineID) == 0 {
					res = append(res, element)
				} 
			}
		resp.Payload = res
	} 
	return resp
}

// JobsList handles the GET endpoints for jobs
// if # args is 0, prints the first 50 jobs on page 1
// if # args is 1, prints the job whose id is the argument
// if @ args is 2, it prints the first argument number of jobs, on the second argument page number
func JobsList(c *cli.Context) error {
	if len(c.Args()) == 0 {
		resp, err := sd.Default.V3.GetV3Jobs(nil)
		if err != nil {
			return cli.ShowSubcommandHelp(c)
		}
		resp = jobsFilterPipeline(resp, c)
		formattedPrint(resp)
	} else if len(c.Args()) == 1 {
		args := c.Args()
		params := v3.NewGetV3JobsIDParams().WithID(args[0])
		resp, err := sd.Default.V3.GetV3JobsID(params)
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		formattedPrint(resp)
	} else if len(c.Args()) == 2{
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
		params := v3.NewGetV3JobsParams().WithCount(&co).WithPage(&p)
		resp, err := sd.Default.V3.GetV3Jobs(params)
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		resp = jobsFilterPipeline(resp, c)
		formattedPrint(resp)
	} else {
		return cli.ShowSubcommandHelp(c)	
	}
	return nil
}
