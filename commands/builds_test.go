package command

import (
	"testing"

	sd "github.com/screwdriver-cd/client/client"
	"github.com/urfave/cli"
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
			t.Error("Expected error to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"builds", "hello", "world"})
}

func TestBuildsGetStep(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsGetStep(sd.Default, c)
		if err == nil {
			t.Error("Expected error to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"builds", "Hello", "World", "Swagger"})
}

func TestBuildsGetStepLogs(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := BuildsGetStepLogs(sd.Default, c)
		if err == nil {
			t.Error("Expected error to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"builds", "abc", "def", "ghi", "jkl"})
}
