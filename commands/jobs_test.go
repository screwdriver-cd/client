package command

import(
	"testing"
	"github.com/urfave/cli"
	// sd "github.com/screwdriver-cd/client/client"
	"github.com/stretchr/testify/assert"
)

func TestJobsList(t *testing.T){
	testApp := cli.NewApp()
	assert := assert.New(t)
	testApp.Action = func(c *cli.Context) error {
		_, err := JobsList(c)
		assert.NotNil(err)
		return nil
	}
	testApp.Run([]string{"jobs", "tomato", "screwdriver", "cd", "cry"})
	testApp.Action = func(c *cli.Context) error {
		_, err := JobsList(c)
		assert.NotNil(err)
		return nil	
	}
	testApp.Run([]string{"jobs", "32", "1x0x0x0"})
}

func TestJobByID(t *testing.T){
	testApp := cli.NewApp()
	assert := assert.New(t)
	testApp.Action = func(c *cli.Context) error {
		_, err := JobByID(c)
		assert.NotNil(err)
		return nil
	}
	testApp.Run([]string{"jobs","a","b","c","d"})
}


