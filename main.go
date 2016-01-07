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

	var (
		j   interface{}
		err error
	)

	for dec.More() {
		err = dec.Decode(&j)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		b, err := json.MarshalIndent(&j, *prefix, *indent)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		fmt.Println(string(b))
	}
	if err != nil {
		os.Exit(1)
	}

}
