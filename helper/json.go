package helper

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(title string, obj interface{}) {
	empJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n%s\n", title, string(empJSON))
}
