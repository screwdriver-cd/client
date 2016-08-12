package command

import(
	"testing"
	"os"

	model "github.com/screwdriver-cd/client/models"
	"github.com/urfave/cli"
	v3 "github.com/screwdriver-cd/client/client/v3"
	sd "github.com/screwdriver-cd/client/client"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/go-swagger/go-swagger/strfmt"
	"time"
)

var testBuild = [] struct {
		Sha string
		Id string
		CreateTime string
		Cause string
		JobId string
		Number float64
		Container string
		Status string
} {
	{
		Sha: "6ab7a23b73310b26f2c6939cd8150a533500a95b",
		Id: "f087bc435e7d2c1cb0a8ec463a348ed0f280b734",
		CreateTime: "2016-08-08T20:55:45.277Z",
		Cause: "Started by user nkatzman",
		JobId: "bc5aa345891e0704972a315d01e93e95afd7ae6a",
		Number: 1470689745277,
		Container: "node:4",
		Status: "QUEUED",
	},
	{
		Sha: "6ab7a23b73310b26f2c6939cd8150a533500a95b",
		Id: "c3a1a3523c94d1f9a8b9a1f5e95f9abb7d612b1a",
		CreateTime: "2016-08-08T21:20:54.319Z",
		Cause: "Started by user nkatzman",
		JobId: "bc5aa345891e0704972a315d01e93e95afd7ae6a",
		Number: 1470691254319,
		Container: "node:4",
		Status: "QUEUED",
	},
	{
		Sha: "4755ebed30459caf74ddcea73f9dca56ef010429",
		Id: "e0a3ba2ec752cbcb45faec57a58aef9eec9d876c",
		CreateTime: "2016-08-06T01:07:05.736Z",
		JobId: "33e8535c1f1efcd52d867272eb8ec2127347e0c2",
		Cause: "Started by user d2lam",
		Number: 1470445625736,
		Container: "node:4",
		Status: "RUNNING",
	},
}


func createMocks() model.ListOfBuilds{

	timeFormat := "2006-01-02T15:04:05Z0700"
	ret := model.ListOfBuilds{}
	for _, element := range testBuild{
			t,_:= time.Parse(timeFormat, element.CreateTime)
			tmpObj := model.GetBuild{
				Cause: element.Cause,
				Container: &element.Container,
				CreateTime: strfmt.Date(t),
				ID: element.Id,
				JobID: element.JobId,
				Sha: &element.Sha,
				Status: element.Status,
			}
			ret = append(ret, &tmpObj)
	}
	return ret
}

func TestBuildsFilterJobs(t *testing.T){
	testApp := cli.NewApp()
	idToFilter := "bc5aa345891e0704972a315d01e93e95afd7ae6a"
	os.Args = []string{"builds","--jobID", idToFilter}
	testApp.Name = "builds"
	testApp.Flags = []cli.Flag{
					cli.StringFlag{Name:"jobID",},	
	}
	gv3BuildsOK := &v3.GetV3BuildsOK{}
	gv3BuildsOK.Payload = createMocks()
	testApp.Action = func(c *cli.Context) error {
		res := buildsFilterJobs(gv3BuildsOK, c)
		for _, element := range res.Payload{
				if element.JobID != idToFilter{
						t.Fail()	
				}
		}	
		return nil
	}
	testApp.Run(os.Args)
}

func TestBuildsFilterStatus(t *testing.T){
	testApp := cli.NewApp()
	statusToFilter := "RUNNING"
	os.Args = []string{"builds", "--status", statusToFilter}
	testApp.Name="builds"
	testApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "status",	
		},
	}
	gv3BuildsOK := &v3.GetV3BuildsOK{}
	gv3BuildsOK.Payload = createMocks()
	testApp.Action = func(c *cli.Context) error {
		res := buildsFilterStatus(gv3BuildsOK, c)
		for _, element := range res.Payload{
				if element.Status != statusToFilter {
					t.Fail()	
				}
		}		
		return nil
	}
	testApp.Run(os.Args)
}

func TestBuildRequestGetBuildList(t *testing.T){
		// httpmock.Activate()
		// defer httpmock.DeactivateAndReset()
    //
		// httpmock.RegisterResponder("get", "http://a4677c9873c9611e6aa7102b92f75d5c-1135862614.us-west-2.elb.amazonaws.com/v3/",
		// httpmock.NewStringResponder(200, `[{"sha": "6ab7a23b73310b26f2c6939cd8150a533500a95b","id": "f087bc435e7d2c1cb0a8ec463a348ed0f280b734","createTime": "2016-08-08T20:55:45.277Z","cause": "Started by user nkatzman","jobId": "bc5aa345891e0704972a315d01e93e95afd7ae6a","number": 1470689745277,"container": "node:4","status": "QUEUED"}]`))
		// 	resp, err := buildRequestGetBuildList(sd.Default)
		// 	if err != nil {
		// 		t.Fail()
		// 	}
		// 	fmt.Println("past fail")
		// 	fmt.Println(len(resp.Payload))
		// 	fmt.Println("yo")
		// 	assert := assert.New(t)
		// 	// assert.Equal(*resp.Payload[0].Sha, "6ab7a23b73310b26f2c6939cd8150a533500a95b", "Equal Sha")
		// 	assert.Equal(resp.Payload[0].ID, "f087bc435e7d2c1cb0a8ec463a348ed0f280b734", "Equal IDs")
		// 	// assert.Equal(resp.Payload[0].CreateTime, "2016-08-08T20:55:45.277Z", "Equal create times")
		// 	// assert.Equal(resp.Payload[0].Status, "QUEUED", "Equal status")
		// 	// assert.Equal(resp.Payload[0].Cause, "Started by user nkatzman", "Assert, started by nkatzman")
		// 	fmt.Println("outofBuildRequest")
}

func TestBuildRequestGetID(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("get", "http://a4677c9873c9611e6aa7102b92f75d5c-1135862614.us-west-2.elb.amazonaws.com/v3/",
	httpmock.NewStringResponder(200, `[{"sha": "6ab7a23b73310b26f2c6939cd8150a533500a95b","id": "f087bc435e7d2c1cb0a8ec463a348ed0f280b734","createTime": "2016-08-08T20:55:45.277Z","cause": "Started by user nkatzman","jobId": "bc5aa345891e0704972a315d01e93e95afd7ae6a","number": 1470689745277,"container": "node:4","status": "QUEUED"}]`))
	resp, err := buildRequestGetID(sd.Default, "f087bc435e7d2c1cb0a8ec463a348ed0f280b734")
	if err != nil {
		t.Fail()	
	}
	assert := assert.New(t)
	assert.Equal(*resp.Payload.Sha, "6ab7a23b73310b26f2c6939cd8150a533500a95b", "Equal Sha")
	assert.Equal(resp.Payload.ID, "f087bc435e7d2c1cb0a8ec463a348ed0f280b734", "Equal IDs")
	assert.Equal(resp.Payload.CreateTime, "2016-08-08T20:55:45.277Z", "Equal create times")
	assert.Equal(resp.Payload.Status, "QUEUED", "Equal status")
	assert.Equal(resp.Payload.Cause, "Started by user nkatzman", "Assert, started by nkatzman")
} 
