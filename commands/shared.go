package command

import(
	"fmt"
	"strconv"
	"encoding/json"
	"errors"

	"github.com/urfave/cli"
)

const(
	IDPARAM = 0
)

const (
	COUNTPARAM int = 0
	PAGENUMPARAM int = 1
	JOBIDPARAM int = 1
)

// formattedPrint Marshals and formats the prints
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

func getCountAndPage(c *cli.Context) (int64, int64, error){
	args := c.Args()
	var count, page int
	var err error
	if len(args) == 2 {
		count, err = strconv.Atoi(args[COUNTPARAM])
		if err != nil {
			return 0, 0, errors.New("Invalid Usage") 
		}
		page, err = strconv.Atoi(args[PAGENUMPARAM])
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
	return args[IDPARAM], nil 
}
