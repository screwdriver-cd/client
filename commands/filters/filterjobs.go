package filter

import(
	"strings"

	"github.com/urfave/cli"
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)

// jobFilterPipeline filters jobs by pipelineID returns the original v3.GetV3JobsOK object
func JobsFilterPipeline(resp *v3.GetV3JobsOK, c *cli.Context) *v3.GetV3JobsOK {
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


