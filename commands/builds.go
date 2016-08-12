package command

import(
	"strconv"
	"strings"

	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)



// buildsFilterJobs returns the passed v3.GetV3BuildsOK object filters out builds whose jobID does not match what was passed in as flag jobID
func buildsFilterJobs(resp *v3.GetV3BuildsOK, c *cli.Context) *v3.GetV3BuildsOK {
	filters := c.String("jobID")
	if strings.Compare(filters, "") != 0 {
			res := model.ListOfBuilds{}
			for _, element := range resp.Payload{ 
				if strings.Compare(element.JobID, filters) == 0 {
					res = append(res, element)
				} 
			}
			resp.Payload = res
	} 
	return resp
}

// buildFilterStatus filters the builds by status that was passed in as flag status. returns the original v3.GetV3BuildsOK object
func buildsFilterStatus(resp *v3.GetV3BuildsOK, c *cli.Context) *v3.GetV3BuildsOK {
	filters := c.String("status")
	if strings.Compare(filters, "") != 0 {
			res := model.ListOfBuilds{}
			for _, element := range resp.Payload{
				if strings.Compare(element.Status, filters) == 0 {
					res = append(res, element)
				} 
			}
			resp.Payload = res
	} 
	return resp
}

// buildRequestGetBuildsListParams takes in a transport, page and count
func buildRequestGetBuildListParams(trans *sd.ScrewdriverAPIDocumentation, count int64, page int64) (*v3.GetV3BuildsOK, error){
	params := v3.NewGetV3BuildsParams().WithCount(&count).WithPage(&page)
	return trans.V3.GetV3Builds(params)
}

// buildRequestGetBuildList takes in a transport, returns the default page (1) and count (50)
func buildRequestGetBuildList(trans *sd.ScrewdriverAPIDocumentation) (*v3.GetV3BuildsOK, error){
	return trans.V3.GetV3Builds(nil)					
}

// BuildList is called by the client command builds-list
// if # args is 0, defaults to listing out 50 builds on page 1
// if # args is 2, gets the first argument number of builds, and the page #
func BuildList(c *cli.Context) error {
	if len(c.Args()) == 0 {
		resp, err := buildRequestGetBuildList(sd.Default)
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		resp = buildsFilterJobs(resp, c)
		resp = buildsFilterStatus(resp, c)
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
		resp, err := buildRequestGetBuildListParams(sd.Default, int64(count), int64(page))
		if err != nil {
			return cli.ShowSubcommandHelp(c)	
		}
		resp = buildsFilterJobs(resp, c)
		resp = buildsFilterStatus(resp, c)
		formattedPrint(resp)
	} else {
		return cli.ShowSubcommandHelp(c)	
	}
	return nil
}

// buildRequestGetID builds the request for getID returns the response and an error
func buildRequestGetID(trans *sd.ScrewdriverAPIDocumentation, id string) (*v3.GetV3BuildsIDOK,error){
		params := v3.NewGetV3BuildsIDParams().WithID(id)
		resp, err := trans.V3.GetV3BuildsID(params)	
		return resp, err
}

// BuildsGetID, given an ID, get the build information
func BuildsGetID(c *cli.Context) error {
		if len(c.Args()) == 1 {
				args := c.Args()	
				resp, err := buildRequestGetID(sd.Default, args[BUILDIDPARAM])
				if err != nil {
						return cli.ShowSubcommandHelp(c)	
				}
				formattedPrint(resp)
		} else {
			return cli.ShowSubcommandHelp(c)	
		}
		return nil
}

