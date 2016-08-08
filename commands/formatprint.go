package command

import(
	"fmt"
	"encoding/json"
)

// formattedPrint Marshals and formats the prints
func formattedPrint(a ...interface{}){
	m, err := json.MarshalIndent(a, " ", "  ")
	if err != nil {
		fmt.Println(err)	
	}
	fmt.Println(string(m))
}
