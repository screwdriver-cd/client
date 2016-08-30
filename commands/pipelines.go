package command

import (
	"errors"

	sd "github.com/screwdriver-cd/client/client"
	v3 "github.com/screwdriver-cd/client/client/v3"
)

// PipelinesList handles the get endpoints for pipelines command
// When number of args are 0, it defaults to getting 50 responses on the first page
// When number of args is 2, the first argument is the number of responses, and the second is the page number
// Prints out information regarding pipelines
func PipelinesList(sdAPIClient *sd.ScrewdriverAPIDocumentation, c Context) (*v3.GetV3PipelinesOK, error) {
	if c.NArg() == 0 {
		return sdAPIClient.V3.GetV3Pipelines(nil)
	} else if c.NArg() == 2 {
		count, page, err := getCountAndPage(c)
		if err != nil {
			return nil, err
		}
		return sdAPIClient.V3.GetV3Pipelines(v3.NewGetV3PipelinesParams().WithCount(&count).WithPage(&page))
	} else {
		return nil, errors.New("Invalid Number of Arguments")
	}
}
