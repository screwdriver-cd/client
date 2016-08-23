package filter

import (
	v3 "github.com/screwdriver-cd/client/client/v3"
	model "github.com/screwdriver-cd/client/models"
	"github.com/urfave/cli"
	"os"
	"testing"
)

func TestBuildsFilterStatus(t *testing.T) {
	testApp := cli.NewApp()
	statusToFilter := "RUNNING"
	os.Args = []string{"builds", "--status", statusToFilter}
	testApp.Name = "builds"
	testApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "status",
		},
	}
	gv3BuildsOK := &v3.GetV3BuildsOK{}
	gv3BuildsOK.Payload = createMocks()
	testApp.Action = func(c *cli.Context) error {
		statusFlag := c.String("status")
		res := BuildsFilterStatus(gv3BuildsOK, statusFlag)
		for _, element := range res.Payload {
			if element.Status != statusToFilter {
				t.Errorf("Expected Statuses not %s to be filtered out", statusToFilter)
			}
		}
		return nil
	}
	testApp.Run(os.Args)
}

func TestBuildsFilterJobs(t *testing.T) {
	testApp := cli.NewApp()
	idToFilter := "bc5aa345891e0704972a315d01e93e95afd7ae6a"
	os.Args = []string{"builds", "--jobID", idToFilter}
	testApp.Name = "builds"
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name: "jobID"},
	}
	gv3BuildsOK := &v3.GetV3BuildsOK{}
	gv3BuildsOK.Payload = createMocks()
	testApp.Action = func(c *cli.Context) error {
		pipelineIDFlag := c.String("pipelineID")
		res := BuildsFilterJobs(gv3BuildsOK, pipelineIDFlag)
		for _, element := range res.Payload {
			if element.JobID != idToFilter {
				t.Errorf("Expected JobIDs not %s to be filtered out", idToFilter)
			}
		}
		return nil
	}
	testApp.Run(os.Args)
}

var testBuild = []struct {
	Sha        string
	ID         string
	CreateTime string
	Cause      string
	JobID      string
	Number     float64
	Container  string
	Status     string
}{
	{
		Sha:        "6ab7a23b73310b26f2c6939cd8150a533500a95b",
		ID:         "f087bc435e7d2c1cb0a8ec463a348ed0f280b734",
		CreateTime: "2016-08-08T20:55:45.277Z",
		Cause:      "Started by user nkatzman",
		JobID:      "bc5aa345891e0704972a315d01e93e95afd7ae6a",
		Number:     1470689745277,
		Container:  "node:4",
		Status:     "QUEUED",
	},
	{
		Sha:        "6ab7a23b73310b26f2c6939cd8150a533500a95b",
		ID:         "c3a1a3523c94d1f9a8b9a1f5e95f9abb7d612b1a",
		CreateTime: "2016-08-08T21:20:54.319Z",
		Cause:      "Started by user nkatzman",
		JobID:      "bc5aa345891e0704972a315d01e93e95afd7ae6a",
		Number:     1470691254319,
		Container:  "node:4",
		Status:     "QUEUED",
	},
	{
		Sha:        "4755ebed30459caf74ddcea73f9dca56ef010429",
		ID:         "e0a3ba2ec752cbcb45faec57a58aef9eec9d876c",
		CreateTime: "2016-08-06T01:07:05.736Z",
		JobID:      "33e8535c1f1efcd52d867272eb8ec2127347e0c2",
		Cause:      "Started by user d2lam",
		Number:     1470445625736,
		Container:  "node:4",
		Status:     "RUNNING",
	},
}

func createMocks() model.ListOfBuilds {
	ret := model.ListOfBuilds{}
	for _, element := range testBuild {
		tmpObj := model.GetBuild{
			Cause:      element.Cause,
			Container:  &element.Container,
			CreateTime: element.CreateTime,
			ID:         element.ID,
			JobID:      element.JobID,
			Sha:        &element.Sha,
			Status:     element.Status,
		}
		ret = append(ret, &tmpObj)
	}
	return ret
}
