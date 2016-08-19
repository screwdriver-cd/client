package command

import(
	"testing"
	"strings"

	"github.com/urfave/cli"
	"github.com/stretchr/testify/assert"
)

//TODO: compare these to outputs
func TestFormattedPrint(t *testing.T){
	assert := assert.New(t)
	err := FormattedPrint("")
	assert.Nil(err)
	FormattedPrint("potato")
	assert.Nil(err)
	FormattedPrint("{{as}")
	assert.Nil(err)
	err = FormattedPrint(func(){})
	assert.Error(err)
}

func TestGetNumArguments(t *testing.T){
	testApp := cli.NewApp()
	assert := assert.New(t)		
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name:"myFlag",},
	}
	testApp.Action = func(c *cli.Context) error {
		assert.Equal(getNumArguments(c), 0)	
		return nil
	}
	testApp.Run([]string{"builds"})
	testApp.Action = func(c *cli.Context) error {
		assert.Equal(getNumArguments(c), 1)
		return nil	
	}
	testApp.Run([]string{"builds", "tomato"})
}

func TestGetCountAndPage(t *testing.T){
	assert := assert.New(t)
	testApp := cli.NewApp()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name:"count",},
		cli.StringFlag{Name:"page",},
	}
	testApp.Action = func(c *cli.Context) error {
			count, page, err := getCountAndPage(c)
			assert.NotNil(err)
			assert.Equal(count, int64(0))
			assert.Equal(page, int64(0))
			return nil
	}
	testApp.Run([]string{"builds"})
	testApp.Action = func(c *cli.Context) error {
			count, page, err := getCountAndPage(c)
			assert.NotNil(err)
			assert.Equal(count, int64(0))
			assert.Equal(page, int64(0))
			return nil	
	}
	testApp.Run([]string{"builds","tomato","swag"})
	testApp.Action = func(c *cli.Context) error {
			count, page, err := getCountAndPage(c)
			assert.Equal(count, int64(50))
			assert.Equal(page, int64(100))
			assert.Nil(err)
			return nil
	}
	testApp.Run([]string{"builds","50","100"})
}

func TestGetID(t *testing.T){
	assert := assert.New(t)
	testApp := cli.NewApp()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name:"ID",},	
	}
	id := "abc123"
	testApp.Action = func(c *cli.Context) error {
			str, err := getID(c)
			assert.Equal(id, str)
			assert.Nil(err)
			return nil
	}
	testApp.Run([]string{"builds", id})
	testApp.Action = func(c *cli.Context) error {
		str, err := getID(c)
		assert.Error(err)
		assert.Equal(strings.Compare(str,""), 0)
		return nil
	}
	testApp.Run([]string{"builds"})
}
