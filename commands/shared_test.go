package command

import (
	"strings"
	"testing"

	"github.com/urfave/cli"
)

func TestFormattedPrint(t *testing.T) {
	err := FormattedPrint("")
	if err != nil {
		t.Fail()
	}
	FormattedPrint("potato")
	FormattedPrint("{{as}")
	err = FormattedPrint(func() {})
	if err == nil {
		t.Fail()
	}
}

func TestGetNumArguments(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name: "myFlag"},
	}
	testApp.Action = func(c *cli.Context) error {
		if getNumArguments(c) != 0 {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds"})
	testApp.Action = func(c *cli.Context) error {
		if getNumArguments(c) != 1 {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds", "tomato"})
}

func TestGetCountAndPage(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name: "count"},
		cli.StringFlag{Name: "page"},
	}
	testApp.Action = func(c *cli.Context) error {
		count, page, err := getCountAndPage(c)
		if err == nil {
			t.Fail()
		}
		if count != int64(0) {
			t.Fail()
		}
		if page != int64(0) {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds"})
	testApp.Action = func(c *cli.Context) error {
		count, page, err := getCountAndPage(c)
		if err == nil {
			t.Fail()
		}
		if count != int64(0) {
			t.Fail()
		}
		if page != int64(0) {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds", "tomato", "swag"})
	testApp.Action = func(c *cli.Context) error {
		count, page, err := getCountAndPage(c)
		if count != int64(50) {
			t.Fail()
		}
		if page != int64(100) {
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds", "50", "100"})
}

func TestGetID(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{Name: "ID"},
	}
	id := "abc123"
	testApp.Action = func(c *cli.Context) error {
		str, err := getID(c)
		if id != str {
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds", id})
	testApp.Action = func(c *cli.Context) error {
		str, err := getID(c)
		if err == nil {
			t.Fail()
		}
		if strings.Compare(str, "") != 0 {
			t.Fail()
		}
		return nil
	}
	testApp.Run([]string{"builds"})
}
