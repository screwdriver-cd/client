package filter

import (
	"strings"

	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)

// JobsFilterPipeline filters jobs by pipelineID returns the original v3.GetV3JobsOK object
func JobsFilterPipeline(resp *v3.GetV3JobsOK, pipelineID string) *v3.GetV3JobsOK {
	res := model.ListOfJobs{}
	for _, element := range resp.Payload {
		if strings.Compare(pipelineID, *element.PipelineID) == 0 {
			res = append(res, element)
		}
	}
	resp.Payload = res
	return resp
}
