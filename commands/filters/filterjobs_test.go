package filter

import (
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
	"github.com/urfave/cli"
	"testing"
)

func TestJobsFilterPipeline(t *testing.T) {
	testApp := cli.NewApp()
	pipelineIDToFilter := "8b723a79b2183a334fe3944197405dc20a3ecb0f"
	gv3JobsOK := &v3.GetV3JobsOK{}
	gv3JobsOK.Payload = createJobMocks()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name: "pipelineID"},
	}
	testApp.Action = func(c *cli.Context) error {
		pipelineIDFlag := c.String("pipelineID")
		res := JobsFilterPipeline(gv3JobsOK, pipelineIDFlag)
		for _, element := range res.Payload {
			if element.PipelineID != pipelineIDToFilter {
				t.Errorf("Expected pipelineIDs not %s", pipelineIDToFilter)
			}
		}
		return nil
	}
}

func createJobMocks() model.ListOfJobs {
	ret := model.ListOfJobs{}
	for _, element := range testJobs {
		tmpObj := model.GetJob{
			ID:         element.ID,
			Name:       element.Name,
			PipelineID: element.PipelineID,
			State:      element.State,
		}
		ret = append(ret, &tmpObj)
	}
	return ret
}

var testJobs = []struct {
	ID         string
	Name       string
	PipelineID string
	State      string
}{
	{
		ID:         "bc5aa345891e0704972a315d01e93e95afd7ae6a",
		Name:       "PR-11",
		PipelineID: "b45729bd24c157b50f19603b60045a1e8b898460",
		State:      "DISABLED",
	},
	{
		ID:         "fbf34b66732f589580c8964314858d0eb9c436ae",
		Name:       "main",
		PipelineID: "8b723a79b2183a334fe3944197405dc20a3ecb0f",
		State:      "ENABLED",
	},
	{
		ID:         "95822ed975ca8dd6856eae54823ff849f7f5b6be",
		Name:       "main",
		PipelineID: "7e1637f07ce250a465595ffc963d5d46b6840e09",
		State:      "ENABLED",
	},
	{
		ID:         "33e8535c1f1efcd52d867272eb8ec2127347e0c2",
		Name:       "PR-9",
		PipelineID: "b45729bd24c157b50f19603b60045a1e8b898460",
		State:      "DISABLED",
	},
	{
		ID:         "ef5e06247c80abf983894839a2d56552016b0c18",
		Name:       "PR-44",
		PipelineID: "4c499806ad2cac5ec98f5cf4805fd3a2bd43203a",
		State:      "DISABLED",
	},
	{
		ID:         "ef09b3cda7fbe9d6134eb7d57f556f80de1bd5b1",
		Name:       "PR-7",
		PipelineID: "7e1637f07ce250a465595ffc963d5d46b6840e09",
		State:      "DISABLED",
	},
	{
		ID:         "e2554d4e44a4a564108ff72f4e265a593feb8fcf",
		Name:       "main",
		PipelineID: "b45729bd24c157b50f19603b60045a1e8b898460",
		State:      "ENABLED",
	},
	{
		ID:         "d42ea144700c978a509183811228763ffd92b9b0",
		Name:       "main",
		PipelineID: "4c499806ad2cac5ec98f5cf4805fd3a2bd43203a",
		State:      "ENABLED",
	},
	{
		ID:         "a079f8d92396091eeb249abe70d3d5441e067b69",
		Name:       "PR-41",
		PipelineID: "4c499806ad2cac5ec98f5cf4805fd3a2bd43203a",
		State:      "ENABLED",
	},
	{
		ID:         "b8ed81caf0e763996b0551a053bffd2d404c7056",
		Name:       "main",
		PipelineID: "b7903fe17af8ec71daf30ae420078db264c7033e",
		State:      "ENABLED",
	},
	{
		ID:         "4c64bacd3840fbe864fe5367654444a1b7d7cc81",
		Name:       "PR-43",
		PipelineID: "4c499806ad2cac5ec98f5cf4805fd3a2bd43203a",
		State:      "DISABLED",
	},
	{
		ID:         "97720302dcd61f5eb53f07b33b56404a465d1a35",
		Name:       "PR-42",
		PipelineID: "4c499806ad2cac5ec98f5cf4805fd3a2bd43203a",
		State:      "DISABLED",
	},
}
