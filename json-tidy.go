// json-tidy pretty prints JSON
package main

import (
	"bytes"
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

	var input, output bytes.Buffer

	input.ReadFrom(os.Stdin)

	die(json.Indent(&output, input.Bytes(), *prefix, *indent))

	output.WriteByte('\n')

	output.WriteTo(os.Stdout)
}
