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
	os.Exit(errors.Execute(Run, nil))
}

func Run(args []string) error {
	fl := flag.NewFlagSet("json-tidy", flag.ContinueOnError)
	prefix := fl.String("prefix", "", "Prefix string")
	indent := fl.String("indent", "\t", "Identation string")
	htmlSafe := fl.Bool("html-safe", false, "Escape special characters for easy embedding in HTML")

	fl.Usage = func() {
		fmt.Fprint(fl.Output(), `Usage of json-tidy:

json-tidy [opts] <file|url|->...
        Gets input (defaults to stdin) and prints clean json to stdout.
`)
		fl.PrintDefaults()
	}
	if err := fl.Parse(args); err != nil {
		return flag.ErrHelp
	}

	args = fl.Args()
	if len(args) == 0 {
		args = []string{flagext.StdIO}
	}
	var errs errors.Slice
	for _, arg := range args {
		errs.Push(tidyPrint(arg, *prefix, *indent, *htmlSafe))
	}
	return errs.Merge()
}

func tidyPrint(arg, prefix, indent string, htmlSafe bool) (err error) {
	src := flagext.FileOrURL(flagext.StdIO, nil)
	if err = src.Set(arg); err != nil {
		return fmt.Errorf("problem with %q: %v\n", arg, err)
	}
	defer errors.Defer(&err, src.Close)

	dec := json.NewDecoder(src)
	dec.UseNumber() // Preserve number formatting

	var data interface{}

	for dec.More() {
		err = dec.Decode(&data)
		if err != nil {
			return fmt.Errorf("problem with %q: %v\n", arg, err)
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent(prefix, indent)
		enc.SetEscapeHTML(htmlSafe)
		err = enc.Encode(&data)
		if err != nil {
			return fmt.Errorf("problem with %q: %v\n", arg, err)
		}
	}
	return nil
}
