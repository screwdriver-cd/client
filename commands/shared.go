package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

// Context interface to context definition by urfave/cli
type Context interface {
	NumFlags() int
	Set(name, value string) error
	GlobalSet(name, value string) error
	IsSet(name string) bool
	GlobalIsSet(name string) bool
	FlagNames() (names []string)
	GlobalFlagNames() (names []string)
	Args() cli.Args
	NArg() int
	String(name string) string
}

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

func getCountAndPage(c Context) (int64, int64, error) {
	args := c.Args()
	var count, page int
	var err error
	if c.NArg() == 2 {
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

func getID(c Context) (string, error) {
	args := c.Args()
	if c.NArg() != 1 {
		return "", errors.New("Invalid number of parameters")
	}
	return args[IDParam], nil
}
