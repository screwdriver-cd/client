package command

import (
	sd "github.com/screwdriver-cd/client/client"
	"github.com/urfave/cli"
	"testing"
)

func TestJobsList(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := JobsList(sd.Default, c)
		if err == nil {
			t.Error("Expected err to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"jobs", "tomato", "screwdriver", "cd", "cry"})
	testApp.Action = func(c *cli.Context) error {
		_, err := JobsList(sd.Default, c)
		if err == nil {
			t.Error("Expected err to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"jobs", "32", "1x0x0x0"})
}

func TestJobByID(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := JobByID(sd.Default, c)
		if err == nil {
			t.Error("Expected err to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"jobs", "a", "b", "c", "d"})
}
