package filter

import(
	"strings"

	"github.com/urfave/cli"
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)

// buildsFilterJobs returns the passed v3.GetV3BuildsOK object filters out builds whose jobID does not match what was passed in as flag jobID
func BuildsFilterJobs(resp *v3.GetV3BuildsOK, c *cli.Context) *v3.GetV3BuildsOK {
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
func BuildsFilterStatus(resp *v3.GetV3BuildsOK, c *cli.Context) *v3.GetV3BuildsOK {
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
