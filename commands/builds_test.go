package command

import (
	sd "github.com/screwdriver-cd/client/client"
	"github.com/urfave/cli"
	"testing"
)

func TestBuildsList(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsList(sd.Default, c)
		if err == nil {
			t.Error("Expected the error to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"builds", "stjohn", "jeeer", "peeetscoffee", "noah", "mintree", "deeren", "dtd", "teef", "sin", "shoooo", "fillz coffee"})
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsList(sd.Default, c)
		if err == nil {
			t.Error("Expect the error to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"builds", "100", "NotANoah"})
}

func TestBuildsGetID(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsGetID(sd.Default, c)
		if err == nil {
			t.Error("Expected error to be nil")
		}
		return nil
	}
	testApp.Run([]string{"jobs", "hello", "world"})
}
