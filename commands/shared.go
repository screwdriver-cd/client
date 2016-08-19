package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

const (
	// IDParam Argument for IDParams (general)
	IDParam = 0
	// CountParam Argument number for count
	CountParam int = 0
	// PageNumParam Argument number for pages
	PageNumParam int = 1

	// JobIDParam Argument number for job ids
	JobIDParam int = 1
)

// FormattedPrint Marshals and formats the prints
func FormattedPrint(a ...interface{}) error {
	m, err := json.MarshalIndent(a, " ", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(m))
	return nil
}

func getNumArguments(c *cli.Context) int {
	return len(c.Args())
}

func getCountAndPage(c *cli.Context) (int64, int64, error) {
	args := c.Args()
	var count, page int
	var err error
	if len(args) == 2 {
		count, err = strconv.Atoi(args[CountParam])
		if err != nil {
			return 0, 0, errors.New("Invalid Usage")
		}
		page, err = strconv.Atoi(args[PageNumParam])
		if err != nil {
			return int64(count), int64(page), errors.New("Invalid USage")
		}

		return int64(count), int64(page), nil
	}
	return 0, 0, errors.New("Invalid Usage")
}

func getID(c *cli.Context) (string, error) {
	args := c.Args()
	if len(args) != 1 {
		return "", errors.New("Invalid number of parameters")
	}
	return args[IDParam], nil
}
