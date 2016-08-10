package command

import(
	// "encoding/json"
	"fmt"
	"testing"
	"os"
	"net/http/httptest"
	"net/http"

	model "github.com/screwdriver-cd/client/models"
	"github.com/urfave/cli"
	v3 "github.com/screwdriver-cd/client/client/v3"
	sd "github.com/screwdriver-cd/client/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/jarcoal/httpmock"
)

type testBuild struct{
		Sha string
		Id string
		CreateTime string
		Cause string
		JobId string
		Number float64
		Container string
		Status string
}

func mockBuildList()[]*testBuild{
	builds := []*testBuild{
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
	return builds
}

func createMocks(arr []*testBuild) model.ListOfBuilds{
	ret := model.ListOfBuilds{}
	for _, element := range arr{
			tmpObj := model.GetBuild{
				Cause: element.Cause,
				Container: &element.Container,
				CreateTime: element.CreateTime,
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
	gv3BuildsOK.Payload = createMocks(mockBuildList())
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
	gv3BuildsOK.Payload = createMocks(mockBuildList())
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
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("hit")
		// mock := createMocks(mockBuildList())
		// json.NewEncoder(w).Encode(mock)
		fmt.Fprint(w,genTest())
	}))
	defer testServer.Close()
	testApp := cli.NewApp()
	os.Args = []string{"builds", "list"}
	testApp.Name = "builds"
	fmt.Println(testServer.URL)
	sd.Default.SetTransport(httptransport.New(testServer.URL[7:], "/", nil))
	testApp.Commands = []cli.Command{
			{
					Name: "list",
					Action: func(c *cli.Context) error {
					BuildList(c)
					return nil
					},
			},	
	}
	testApp.Run(os.Args)
}

func genTest() string{
				TEST2 := `{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"f087bc435e7d2c1cb0a8ec463a348ed0f280b734","createTime":"2016-08-08T20:55:45.277Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470689745277,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"c3a1a3523c94d1f9a8b9a1f5e95f9abb7d612b1a","createTime":"2016-08-08T21:20:54.319Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470691254319,"container":"node:4","status":"QUEUED"},{"sha":"4755ebed30459caf74ddcea73f9dca56ef010429","id":"e0a3ba2ec752cbcb45faec57a58aef9eec9d876c","createTime":"2016-08-06T01:07:05.736Z","jobId":"33e8535c1f1efcd52d867272eb8ec2127347e0c2","cause":"Started by user d2lam","number":1470445625736,"container":"node:4","status":"RUNNING"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"1bfa36a49cad922561d72fd6219d51b795901356","createTime":"2016-08-06T01:19:47.369Z","cause":"Started by user d2lam","jobId":"33e8535c1f1efcd52d867272eb8ec2127347e0c2","number":1470446387369,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"6adb6468472e0b06bf4c54701e044725c1147d2b","createTime":"2016-08-08T20:54:08.101Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470689648101,"container":"node:4","status":"QUEUED"},{"sha":"a137f8b3db6f8536778f89a2953bb445650736e6","id":"de52aa924c6860442e2dd9149c633650ac913ff0","createTime":"2016-08-08T23:04:18.473Z","cause":"Started by user stjohnjohnson","jobId":"ef09b3cda7fbe9d6134eb7d57f556f80de1bd5b1","number":1470697458473,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"f80efc23c01aa039dfc1212e58be7b74f1ccd77d","createTime":"2016-08-06T01:05:53.791Z","cause":"Started by user d2lam",`
TEST2+=`"jobId":"e2554d4e44a4a564108ff72f4e265a593feb8fcf","number":1470445553791,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"e95ce5888bd495bc1568a0e85670083b32dbfaf8","createTime":"2016-08-08T21:21:20.520Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470691280520,"container":"node:4","status":"QUEUED"},{"sha":"d8b8ed6a40beb48e797bb342bdfb8f100375e719","id":"449841687f6e9b98e1781787b818c0f12fa47452","createTime":"2016-08-08T20:50:48.511Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470689448511,"container":"node:4","status":"QUEUED"},{"sha":"d8b8ed6a40beb48e797bb342bdfb8f100375e719","id":"feb582d0db44e9a79d32c8829b6b11e292d8bd09","createTime":"2016-08-08T20:50:48.573Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470689448573,"container":"node:4","status":"QUEUED"},{"sha":"a137f8b3db6f8536778f89a2953bb445650736e6","id":"3316cddfce65fad19fa31eafb64e9c868e49f81a","createTime":"2016-08-08T23:01:39.586Z","cause":"Started by user stjohnjohnson","jobId":"ef09b3cda7fbe9d6134eb7d57f556f80de1bd5b1","number":1470697299586,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"a2a0ae164ad2725ac001e59eb17333b54d9e5e16","createTime":"2016-08-08T21:20:54.324Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470691254324,"container":"node:4","status":"QUEUED"},{"sha":"a137f8b3db6f8536778f89a2953bb445650736e6","id":"518d5a08690df0cbdc928e23e166161b26ffa499","createTime":"2016-08-08T21:31:43.401Z","cause":"Started by user stjohnjohnson","jobId":"95822ed975ca8dd6856eae54823ff849f7f5b6be","number":1470691903401,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"c1f756e96369f8e50f29f52da0d99b126bf711f2","createTime":"2016-08-08T20:54:08.064Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470689648064,"container":"node:4","status":"QUEUED"},{"sha":"b7a67d6f00edf26ecc82ce306abc287b6ec6861b","id":"6c495aa170429da7790046eb1cf8792242edd838","createTime":"2016-08-08T22:27:36.205Z","cause":"Started by user stjohnjohnson","jobId":"ef09b3cda7fbe9d6134eb7d57f556f80de1bd5b1","number":1470695256205,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"7047434081f061156e967c403eb49f2cd0ebfad4","createTime":"2016-08-08T20:55:45.274Z","cause":"Started by user nkatzman","jobId":"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470689745274,"container":"node:4","status":"QUEUED"},{"sha":"6ab7a23b73310b26f2c6939cd8150a533500a95b","id":"c61fe813a8dd7da5bdcddb7ae34a70381c8ce211","createTime":"2016-08-08T21:21:20.527Z","cause":"Started by user nkatzman","jobId":`
TEST2+=`"bc5aa345891e0704972a315d01e93e95afd7ae6a","number":1470691280527,"container":"node:4","status":"QUEUED"},{"sha":"a137f8b3db6f8536778f89a2953bb445650736e6","id":"76769810d2b199cff689b6f7ece0afeb82355fd8","createTime":"2016-08-08T22:28:30.701Z","cause":"Started by user stjohnjohnson","jobId":"ef09b3cda7fbe9d6134eb7d57f556f80de1bd5b1","number":1470695310701,"container":"node:4","status":"QUEUED"}`
return TEST2
}
