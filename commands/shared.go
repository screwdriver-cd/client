package command

import(
	"fmt"
	"strconv"
	"encoding/json"
	"errors"

	"github.com/urfave/cli"
)

const(
	IDPARAM = 1
)

const (
	COUNTPARAM int = 0
	PAGENUMPARAM int = 1
	JOBIDPARAM int = 1
)

// formattedPrint Marshals and formats the prints
func formattedPrint(a ...interface{}){
	m, err := json.MarshalIndent(a, " ", "  ")
	if err != nil {
		fmt.Println(err)	
	}
	fmt.Println(string(m))
}

func getNumArguments(c *cli.Context) int {
	return len(c.Args())
}

func getCountAndPage(c *cli.Context) (int, int, error){
	args := c.Args()
	if len(args) == 2 {
		count, err := strconv.Atoi(args[COUNTPARAM])
		page, err := strconv.Atoi(args[PAGENUMPARAM])
		if err != nil {
			return 0, 0, err	
		}
		return count, page, nil
	}
	return 0,0,errors.New("Invalid number of arguments")
}

func getID(c *cli.Context) (string, error) {
	args := c.Args()
	if len(args) != 1{
		return "", errors.New("Invalid number of parameters")	
	}
	return args[IDPARAM], nil 
}
