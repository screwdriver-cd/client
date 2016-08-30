package command

import (
	"errors"

	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
)

// BuildsList is called by the client command builds-list
// if # args is 0, defaults to listing out 50 builds on page 1
// if # args is 2, gets the first argument number of builds, and the page #
func BuildsList(sdAPIClient *sd.ScrewdriverAPIDocumentation, c Context) (*v3.GetV3BuildsOK, error) {
	if c.NArg() == 0 {
		return sdAPIClient.V3.GetV3Builds(nil)
	} else if c.NArg() == 2 {
		count, page, err := getCountAndPage(c)
		if err != nil {
			return nil, err
		}
		return sdAPIClient.V3.GetV3Builds(v3.NewGetV3BuildsParams().WithCount(&count).WithPage(&page))
	}
	return nil, errors.New("Invalid Number of Arguments")
}

// BuildsGetID given an ID, get the build information
func BuildsGetID(sdAPIClient *sd.ScrewdriverAPIDocumentation, c Context) (*v3.GetV3BuildsIDOK, error) {
	id, err := getID(c)
	if err != nil {
		return nil, err
	}
	return sdAPIClient.V3.GetV3BuildsID(v3.NewGetV3BuildsIDParams().WithID(id))
}
