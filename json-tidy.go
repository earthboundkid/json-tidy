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
	prefix   = flag.String("prefix", "", "Prefix string")
	indent   = flag.String("indent", "\t", "Identation string")
	htmlSafe = flag.Bool("html-safe", false, "Escape special characters for easy embedding in HTML")
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run() (err error) {
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
		return errors.New("Too many arguments")
	}

	var src io.Reader = os.Stdin
	if arg := flag.Arg(0); arg != "" && arg != "-" {
		if u, err := url.Parse(arg); err == nil &&
			// It's a URL
			u.Scheme == "http" || u.Scheme == "https" {
			rsp, err := http.Get(arg)
			if err != nil {
				return err
			}
			defer DeferClose(&err, rsp.Body.Close)
			src = rsp.Body
		} else {
			// It's a file
			f, err := os.Open(arg)
			if err != nil {
				return err
			}
			defer DeferClose(&err, f.Close)
			src = f
		}
	}

	dec := json.NewDecoder(src)
	dec.UseNumber() // Preserve number formatting

	var data interface{}

	for dec.More() {
		err = dec.Decode(&data)
		if err != nil {
			return err
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent(*prefix, *indent)
		enc.SetEscapeHTML(*htmlSafe)
		err = enc.Encode(&data)
		if err != nil {
			return err
		}
	}
	// Trailing newline
	_, err = os.Stdout.Write([]byte{'\n'})
	return err
}

func DeferClose(err *error, f func() error) {
	newErr := f()
	if *err == nil {
		*err = newErr
	}
}
