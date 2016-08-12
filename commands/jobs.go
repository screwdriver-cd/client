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

func buildRequestGetJobList(trans *sd.ScrewdriverAPIDocumentation) (*v3.GetV3JobsOK, error){
	return trans.V3.GetV3Jobs(nil)	
}

func buildRequestGetJobListParams(trans *sd.ScrewdriverAPIDocumentation, count int64, page int64)(*v3.GetV3JobsOK, error){
	params := v3.NewGetV3JobsParams().WithCount(&count).WithPage(&page)
	return trans.V3.GetV3Jobs(params)
}

// JobsList handles the GET endpoints for jobs
// if # args is 0, prints the first 50 jobs on page 1
// if # args is 2, it prints the first argument number of jobs, on the second argument page number
func JobsList(c *cli.Context) error {
	if len(c.Args()) == 0 {
		resp, err := buildRequestGetJobList(sd.Default)
		if err != nil {
			return cli.ShowSubcommandHelp(c)
		}
		resp = jobsFilterPipeline(resp, c)
		formattedPrint(resp)
	} else if len(c.Args()) == 2{
		args := c.Args()	
		count, err := strconv.Atoi(args[COUNTPARAM])
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		page, err := strconv.Atoi(args[PAGENUMPARAM])
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		resp, err := buildRequestGetJobListParams(sd.Default, int64(count), int64(page))
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


func buildRequestGetJobID(trans *sd.ScrewdriverAPIDocumentation, id string) (*v3.GetV3JobsIDOK, error){
	params := v3.NewGetV3JobsIDParams().WithID(id)
	return sd.Default.V3.GetV3JobsID(params)
}

// JobByID Get a specified job by ID returns an error if unable to marshal data or unable to connect
func JobByID(c *cli.Context) error {
	if len(c.Args()) == 1 {
		args := c.Args()
		resp, err := buildRequestGetJobID(sd.Default, args[JOBIDPARAM])
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		formattedPrint(resp)
	} else {
		return cli.ShowSubcommandHelp(c)	
	}
	return nil
}

