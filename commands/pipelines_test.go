package command

import (
	"testing"

	sd "github.com/screwdriver-cd/client/client"
	"github.com/urfave/cli"
)

func TestPipelinesList(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := PipelinesList(sd.Default, c)
		if err == nil {
			t.Error("Expected err to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"yolo", "swagger", "fly", "beats", ":)", "swiggity"})
	testApp.Action = func(c *cli.Context) error {
		_, err := PipelinesList(sd.Default, c)
		if err == nil {
			t.Error("Expected err to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"yolo", "52", "1XOXO"})
}

func TestPipelinesGetID(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Action = func(c *cli.Context) error {
		_, err := PipelinesGetID(sd.Default, c)
		if err == nil {
			t.Error("Expected err to be not nil")
		}
		return nil
	}
	testApp.Run([]string{"yolo"})
}
