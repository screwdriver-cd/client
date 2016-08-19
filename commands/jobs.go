package command

import(
	"errors"

	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
)

// JobsList handles the GET endpoints for jobs
// if # args is 0, prints the first 50 jobs on page 1
// if # args is 2, it prints the first argument number of jobs, on the second argument page number
func JobsList(c *cli.Context) (*v3.GetV3JobsOK, error) {
	numParams := getNumArguments(c)
	if numParams == 0 {
		return sd.Default.V3.GetV3Jobs(nil)
	} else if numParams == 2{
		count, page, err := getCountAndPage(c)	
		if err != nil {
			return nil, err	
		}
		return sd.Default.V3.GetV3Jobs(v3.NewGetV3JobsParams().WithCount(&count).WithPage(&page))
	}
	return nil, errors.New("Invalid Usage")
}

// JobByID Get a specified job by ID returns an error if unable to marshal data or unable to connect
func JobByID(c *cli.Context) (*v3.GetV3JobsIDOK, error) {
	id, err := getID(c)
	if err != nil {
		return nil, err	
	}
	return sd.Default.V3.GetV3JobsID(v3.NewGetV3JobsIDParams().WithID(id))
}

