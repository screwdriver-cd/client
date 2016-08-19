package command

import(
	"testing"
	"github.com/urfave/cli"
	// model "github.com/screwdriver-cd/client/models"
	// v3 "github.com/screwdriver-cd/client/client/v3"
	// sd "github.com/screwdriver-cd/client/client"
	"github.com/stretchr/testify/assert"
)

func TestBuildsList(t *testing.T){
	testApp :=cli.NewApp()
	assert := assert.New(t)
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsList(c)
		assert.NotNil(err)
		return nil
	}
	testApp.Run([]string{"builds", "stjohn","jeeer","peeetscoffee","noah","mintree","deeren","dtd","teef","sin","shoooo","fillz coffee"})
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsList(c)
		assert.NotNil(err)
		return nil
	}	
	testApp.Run([]string{"builds", "100", "NotANoah"})
}

func TestBuildsGetID(t *testing.T){
	testApp := cli.NewApp()
	assert := assert.New(t)
	testApp.Action = func(c *cli.Context) error {
			_, err := BuildsGetID(c)
			assert.NotNil(err)
			return nil
	}
	testApp.Run([]string{"jobs", "hello", "world"})
}
