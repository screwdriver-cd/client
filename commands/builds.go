package command

import(
	"errors"

	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
)

// BuildList is called by the client command builds-list
// if # args is 0, defaults to listing out 50 builds on page 1
// if # args is 2, gets the first argument number of builds, and the page #
func BuildsList(c *cli.Context) (*v3.GetV3BuildsOK, error) {
	numParams := getNumArguments(c)
	if numParams == 0{
		return sd.Default.V3.GetV3Builds(nil)
	} else if numParams == 2{
		count, page, err := getCountAndPage(c)
		if err != nil {
			return nil, err	
		}
		return sd.Default.V3.GetV3Builds(v3.NewGetV3BuildsParams().WithCount(&count).WithPage(&page))
	}
	return nil, errors.New("Invalid Number of Arguments")
}

// BuildsGetID, given an ID, get the build information
func BuildsGetID(c *cli.Context) (*v3.GetV3BuildsIDOK, error) {
		id, err := getID(c)
		if err != nil {
			return nil, err	
		}
		return sd.Default.V3.GetV3BuildsID(v3.NewGetV3BuildsIDParams().WithID(id))
}
