package command

import(
	"strconv"
	"strings"
	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)

func buildsFilterJobs(resp *v3.GetV3BuildsOK, c *cli.Context) *v3.GetV3BuildsOK {
	filters := c.String("jobID")
	if strings.Compare(filters, "") != 0 {
			res := model.ListOfBuilds{}
			for _, element := range resp.Payload{ //adding to the same thing
				if strings.Compare(element.JobID, filters) == 0 {
					res = append(res, element)
				} 
			}
			resp.Payload = res
	} 
	return resp
}

func buildsFilterStatus(resp *v3.GetV3BuildsOK, c *cli.Context) *v3.GetV3BuildsOK {
	filters := c.String("status")
	if strings.Compare(filters, "") != 0 {
			res := model.ListOfBuilds{}
			for _, element := range resp.Payload{
				if strings.Compare(element.Status, filters) != 0 {
					res = append(res, element)
				} 
			}
			resp.Payload = res
	} 
	return resp
}

// Builds List endpoint
func BuildList(c *cli.Context) error {
	if len(c.Args()) == 0 {
		resp, err := sd.Default.V3.GetV3Builds(nil)
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		resp = buildsFilterJobs(resp, c)
		resp = buildsFilterStatus(resp, c)
		formattedPrint(resp)
	} else if len(c.Args()) == 1{
		args := c.Args()	
		params := v3.NewGetV3BuildsIDParams().WithID(args[0])
		resp, err := sd.Default.V3.GetV3BuildsID(params)
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
		params := v3.NewGetV3BuildsParams().WithCount(&p).WithPage(&co)
		resp, err := sd.Default.V3.GetV3Builds(params)
		resp = buildsFilterJobs(resp, c)
		resp = buildsFilterStatus(resp, c)
		formattedPrint(resp)
	} else {
		return cli.ShowSubcommandHelp(c)	
	}
	return nil
}
