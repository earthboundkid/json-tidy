// json-tidy pretty prints JSON
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	prefix = flag.String("prefix", "", "Prefix string")
	indent = flag.String("indent", "\t", "Identation string")
)

func main() {
	flag.Parse()

	dec := json.NewDecoder(os.Stdin)

	var j interface{}

	for dec.More() {
		err := dec.Decode(&j)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		b, err := json.MarshalIndent(&j, *prefix, *indent)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	}
}
