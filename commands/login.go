package command

import(
	"fmt"
	sd "github.com/screwdriver-cd/client/client"
	// v3 "github.com/screwdriver-cd/client/client/v3"
)

func Login() error {
	fmt.Println("Please visit the following url in order to login and generate your token")	
	fmt.Println(sd.Default.V3.GetV3Login(nil))
	return sd.Default.V3.GetV3Login(nil)
}
