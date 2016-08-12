package command

import(
	"testing"
	"github.com/jarcoal/httpmock"
	sd "github.com/screwdriver-cd/client/client"
	"github.com/stretchr/testify/assert"
)


func TestBuildRequestGetPipelinesList(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("get", "http://a4677c9873c9611e6aa7102b92f75d5c-1135862614.us-west-2.elb.amazonaws.com/v3/",
		httpmock.NewStringResponder(200, `[ {
         "admins": {
           "d2lam": true
         },
         "configUrl": "git@github.com:screwdriver-cd/models.git#newformat",
         "createTime": "2016-08-06T01:18:54.082Z",
         "id": "b7903fe17af8ec71daf30ae420078db264c7033e",
         "scmUrl": "git@github.com:screwdriver-cd/models.git#newformat"
       },
       {
         "admins": {
           "stjohnjohnson": true
         },
         "configUrl": "git@github.com:screwdriver-cd/config-parser.git#master",
         "createTime": "2016-08-08T21:31:19.596Z",
         "id": "7e1637f07ce250a465595ffc963d5d46b6840e09",
         "scmUrl": "git@github.com:screwdriver-cd/config-parser.git#master"
       },
       {
         "admins": {
           "d2lam": true
         },
         "configUrl": "git@github.com:screwdriver-cd/hashr.git#master",
         "createTime": "2016-08-06T01:05:29.985Z",
         "id": "b45729bd24c157b50f19603b60045a1e8b898460",
         "scmUrl": "git@github.com:screwdriver-cd/hashr.git#master"
       },
       {
         "admins": {
           "tkyi": true
         },
         "configUrl": "git@github.com:screwdriver-cd/models.git#master",
         "createTime": "2016-08-09T19:55:47.227Z",
         "id": "4c499806ad2cac5ec98f5cf4805fd3a2bd43203a",
         "scmUrl": "git@github.com:screwdriver-cd/models.git#master"
       },
       {
         "admins": {
           "tkyi": true
         },
         "configUrl": "git@github.com:screwdriver-cd/models.git#MASter",
         "createTime": "2016-08-09T19:55:52.499Z",
         "id": "8b723a79b2183a334fe3944197405dc20a3ecb0f",
         "scmUrl": "git@github.com:screwdriver-cd/models.git#MASter"
       }
     ]`))
	resp, err := buildRequestGetPipelines(sd.Default)
	if err != nil {
		t.Fail()	
	}
	assert := assert.New(t)
	assert.Equal(resp.Payload[0].ID, "b7903fe17af8ec71daf30ae420078db264c7033e", "EqualIDs")
	assert.Equal(resp.Payload[0].ScmURL, "git@github.com:screwdriver-cd/models.git#newformat", "scmURL assert equal")
	assert.Equal(resp.Payload[0].CreateTime, "2016-08-06T01:18:54.082Z", "Assert creation times are equal")
}
