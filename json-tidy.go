// json-tidy pretty prints JSON
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage of json-tidy:

json-tidy [opts] [file|url|-]
        Gets input (defaults to stdin) and prints clean json to stdout.
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		die(errors.New("Too many arguments"))
	}

	var src io.Reader = os.Stdin
	if arg := flag.Arg(0); arg != "" && arg != "-" {
		if u, err := url.Parse(arg); err == nil &&
			// It's a URL
			u.Scheme == "http" || u.Scheme == "https" {
			rsp, err := http.Get(arg)
			die(err)
			defer rsp.Body.Close()
			src = rsp.Body
		} else {
			// It's a file
			f, err := os.Open(arg)
			die(err)
			defer f.Close()
			src = f
		}
	}

	dec := json.NewDecoder(src)
	dec.UseNumber() // Preserve number formatting

	var data interface{}

	for dec.More() {
		die(dec.Decode(&data))

		b, err := json.MarshalIndent(&data, *prefix, *indent)
		die(err)

		_, err = os.Stdout.Write(b)
		die(err)
	}
	// Trailing newline
	os.Stdout.Write([]byte{'\n'})
}
