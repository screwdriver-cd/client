package command

import (
	"errors"

	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
	"github.com/urfave/cli"
)

// BuildsList is called by the client command builds-list
// if # args is 0, defaults to listing out 50 builds on page 1
// if # args is 2, gets the first argument number of builds, and the page #
func BuildsList(sdAPIClient *sd.ScrewdriverAPIDocumentation, c *cli.Context) (*v3.GetV3BuildsOK, error) {
	numParams := getNumArguments(c)
	if numParams == 0 {
		return sdAPIClient.V3.GetV3Builds(nil)
	} else if numParams == 2 {
		count, page, err := getCountAndPage(c)
		if err != nil {
			return nil, err
		}
		return sdAPIClient.V3.GetV3Builds(v3.NewGetV3BuildsParams().WithCount(&count).WithPage(&page))
	}
	return nil, errors.New("Invalid Number of Arguments")
}

// BuildsGetID given an ID, get the build information
func BuildsGetID(sdAPIClient *sd.ScrewdriverAPIDocumentation, c *cli.Context) (*v3.GetV3BuildsIDOK, error) {
	id, err := getID(c)
	if err != nil {
		return nil, err
	}
	return sdAPIClient.V3.GetV3BuildsID(v3.NewGetV3BuildsIDParams().WithID(id))
}

// BuildsGetStep given an ID and stepName, get the step record
func BuildsGetStep(sdAPIClient *sd.ScrewdriverAPIDocumentation, c *cli.Context) (*v3.GetV3BuildsIDStepsNameOK, error) {
	id, stepName, err := getIDAndStep(c)
	if err != nil {
		return nil, err
	}
	return sdAPIClient.V3.GetV3BuildsIDStepsName(v3.NewGetV3BuildsIDStepsNameParams().WithID(id).WithName(stepName))
}

// BuildsGetStepLogs given an ID and stepName, with optional flag to be set to specifiy which log number to start
func BuildsGetStepLogs(sdAPIClient *sd.ScrewdriverAPIDocumentation, c *cli.Context) (*v3.GetV3BuildsIDStepsNameLogsOK, error) {
	id, stepName, err := getIDAndStep(c)
	if err != nil {
		return nil, err
	}
	var from int64 = int64(c.Int("start"))
	return sdAPIClient.V3.GetV3BuildsIDStepsNameLogs(v3.NewGetV3BuildsIDStepsNameLogsParams().WithID(id).WithName(stepName).WithFrom(&from))
}
