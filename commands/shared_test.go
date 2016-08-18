package command

import(
	"testing"
	"os"
	// "strings"

	"github.com/urfave/cli"
	"github.com/stretchr/testify/assert"
)

//TODO: compare these to outputs
func TestFormattedPrint(t *testing.T){
	formattedPrint("")
	formattedPrint("potato")
	formattedPrint("{{as}")
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
	testApp.Run(os.Args)
	os.Args = []string{"builds", "tomato"}
	testApp.Action = func(c *cli.Context) error {
		assert.Equal(getNumArguments(c), 1)
		return nil	
	}
	testApp.Run(os.Args)
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
			assert.Equal(count, 0)
			assert.Equal(page, 0)
			return nil
	}
	testApp.Run(os.Args)
	os.Args = []string{"builds", "--count", "tomato", "--page", "swag"}
	testApp.Action = func(c *cli.Context) error {
			count, page, err := getCountAndPage(c)
			assert.NotNil(err)
			assert.Equal(count, 0)
			assert.Equal(page, 0)
			return nil	
	}
	testApp.Run(os.Args)
	os.Args = []string{"builds", "50", "100"}
	testApp.Action = func(c *cli.Context) error {
			count, page, err := getCountAndPage(c)
			assert.Equal(count, 50)
			assert.Equal(page, 100)
			assert.Nil(err)
			return nil
	}
	testApp.Run(os.Args)
}
//
// func TestGetID(t *testing.T){
// 	assert := assert.New(t)
// 	testApp := cli.NewApp()
// 	testApp.Flags = []cli.Flag{
// 		cli.StringFlag{Name:"ID",},	
// 	}
// 	id := "abc123"
// 	os.Args =[]string{"--ID", id}
// 	testApp.Action = func(c *cli.Context) error {
// 			str, err := getID(c)
// 			assert.Equal(id, str)
// 			assert.Nil(err)
// 			return nil
// 	}
// 	testApp.Run(os.Args)
// 	testApp.Action = func(c *cli.Context) error {
// 		str, err := getID(c)
// 		assert.Error(err)
// 		assert.Equal(strings.Compare(str,""), 0)
// 		return nil
// 	}
// 	os.Args = []string{}
// 	testApp.Run(os.Args)
// }
