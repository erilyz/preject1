// fmt.Println("usage: go run <appName.go>  < json/<filename>.json")
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

type passedInDataStruct struct {
	Name string `json:"name,omitempty"`
}

func main() {

	flag.Parse()
	var data []byte
	var name string
	var err error

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// read json file specified in command line
		data, err = ioutil.ReadAll(os.Stdin)

		if err != nil {
			name = "World"
		} else {
			// unmarshal json from file to struct
			var temp passedInDataStruct
			err := json.Unmarshal(data, &temp)
			if err != nil {
				os.Exit(1)
			}

			name = temp.Name
		}

	} else {
		name = "World"
	}

	println("Hello", name)
}
