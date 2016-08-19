package command

import(
	// "errors"

	"github.com/urfave/cli"
	sd "github.com/screwdriver-cd/client/client"
	// v3 "github.com/screwdriver-cd/client/client/v3"
)


func Stats(c *cli.Context)  error{
	return sd.Default.V3.GetV3Stats(nil)
}

func Status(c *cli.Context)  error {
	return sd.Default.V3.GetV3Status(nil)
}
