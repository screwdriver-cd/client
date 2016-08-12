package command

import(
	"fmt"
	"encoding/json"
)

const(
	BUILDIDPARAM = 1
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
