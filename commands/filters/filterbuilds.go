package filter

import (
	"strings"

	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
)

// BuildsFilterJobs returns the passed v3.GetV3BuildsOK object filters out builds whose jobID does not match what was passed in as flag jobID
func BuildsFilterJobs(resp *v3.GetV3BuildsOK, jobIDFilter string) *v3.GetV3BuildsOK {
	res := model.ListOfBuilds{}
	for _, element := range resp.Payload {
		asStr := element.JobID
		if strings.Compare(*asStr, jobIDFilter) == 0 {
			res = append(res, element)
		}
	}
	resp.Payload = res
	return resp
}

// BuildsFilterStatus filters the builds by status that was passed in as flag status. returns the original v3.GetV3BuildsOK object
func BuildsFilterStatus(resp *v3.GetV3BuildsOK, statusFilter string) *v3.GetV3BuildsOK {
	res := model.ListOfBuilds{}
	for _, element := range resp.Payload {
		asStr := element.Status
		if strings.Compare(*asStr, statusFilter) == 0 {
			res = append(res, element)
		}
	}
	resp.Payload = res
	return resp
}
