// json-tidy pretty prints JSON
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/carlmjohnson/errors"
	"github.com/carlmjohnson/flagext"
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run() (err error) {
	prefix := flag.String("prefix", "", "Prefix string")
	indent := flag.String("indent", "\t", "Identation string")
	htmlSafe := flag.Bool("html-safe", false, "Escape special characters for easy embedding in HTML")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage of json-tidy:

json-tidy [opts] <file|url|->...
        Gets input (defaults to stdin) and prints clean json to stdout.
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{""}
	}
	for _, arg := range args {
		src := flagext.FileOrURL(flagext.StdIO, nil)
		if err = src.Set(arg); err != nil {
			return err
		}
		defer errors.Defer(&err, src.Close)

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
	}
	return nil
}
