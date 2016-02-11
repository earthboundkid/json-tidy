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

func die(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	dec := json.NewDecoder(os.Stdin)

	var data interface{}

	for dec.More() {
		die(dec.Decode(&data))

		b, err := json.MarshalIndent(&data, *prefix, *indent)
		die(err)

		_, err = os.Stdout.Write(b)
		die(err)
	}
}
