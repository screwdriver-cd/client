package command

import(
// 	"fmt"
	"testing"
// 	"net/http/httptest"
// 	"net/http"
// 	sd "github.com/screwdriver-cd/client/client"
// 	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
// 	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/urfave/cli"
	"github.com/stretchr/testify/assert"
	// "os"
)

// cache-control:no-cache
// Connection:keep-alive
// content-encoding:gzip
// content-type:application/json; charset=utf-8
// Date:Thu, 18 Aug 2016 18:55:32 GMT
// Transfer-Encoding:chunked
// vary:accept-encoding
//
// var pipelinesRes = `[{"id":"6a17984b68cf96616352db5bba422fd46f94564d","scmUrl":"git@github.com:screwdriver-cd/client.git#master","configUrl":"git@github.com:screwdriver-cd/client.git#master","createTime":"2016-08-15T17:32:18.154Z","admins":{"tkyi":true}},{"id":"4c499806ad2cac5ec98f5cf4805fd3a2bd43203a","scmUrl":"git@github.com:screwdriver-cd/models.git#master","configUrl":"git@github.com:screwdriver-cd/models.git#master","createTime":"2016-08-12T22:54:11.406Z","admins":{"tkyi":true}},{"id":"a5dd3581ad1495d758a55abbb8ba6a1349a1e6ed","scmUrl":"git@github.com:screwdriver-cd/gitversion.git#master","configUrl":"git@github.com:screwdriver-cd/gitversion.git#master","createTime":"2016-08-15T17:33:57.467Z","admins":{"tkyi":true}}]`
// func TestBuildRequestGetPipelines(t *testing.T){
// 				println("starting test to request get pipelines")
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 		w.WriteHeader(200)
// 		w.Header().Set("content-type", "application/json; charset=utf-8")
// 		w.Header().Set("Transfer-Encoding", "chunked")
// 		w.Header().Set("content-encoding", "gzip")
// 		w.Header().Set("Connection", "keep-alive")
// 		fmt.Fprint(w, pipelinesRes)
// 	}))
// 	defer server.Close()
//
// 	transport := httptransport.New(server.URL[7:], "/", []string{"http"})
// 	cli := sd.New(transport, strfmt.Default)
// 	_, err := buildRequestGetPipelines(cli)
// 	if err != nil {
// 		println(":sadface")
// 		fmt.Println(err)
// 	}
//
// }
func TestPipelinesList(t *testing.T){
	testApp := cli.NewApp()
	assert := assert.New(t)
	testApp.Action = func(c *cli.Context) error {
		_, err := PipelinesList(c)	
		assert.NotNil(err)
		return nil			
	}
	testApp.Run([]string{"yolo","swagger", "fly", "beats", ":)", "swiggity"})
	testApp.Action = func(c *cli.Context) error {
		_, err := PipelinesList(c)
		assert.NotNil(err)
		return nil
	}
	testApp.Run([]string{"yolo", "52", "1XOXO"})
	// testApp.Action = func(c *cli.Context) error {
	// 	err := PipelinesList(c)	
	// 	assert.Nil(err)
	// 	return nil
	// }
	// testApp.Run([]string{"yolo", "50", "100"})
}
