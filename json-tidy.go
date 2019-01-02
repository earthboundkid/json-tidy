// json-tidy pretty prints JSON
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/carlmjohnson/errors"
	"github.com/carlmjohnson/flagext"
)

func main() {
	os.Exit(errors.Execute(Run, nil))
}

func Run(args []string) (err error) {
	fl := flag.NewFlagSet("json-tidy", flag.ContinueOnError)
	prefix := fl.String("prefix", "", "Prefix string")
	indent := fl.String("indent", "\t", "Identation string")
	htmlSafe := fl.Bool("html-safe", false, "Escape special characters for easy embedding in HTML")
	dst := flagext.FileWriter(flagext.StdIO)
	fl.Var(dst, "output", "write tidy JSON to `file`")
	fl.Usage = func() {
		fmt.Fprint(fl.Output(), `Gets input files and URLs (defaults to stdin) and outputs tidy JSON.

Usage of json-tidy:

json-tidy [opts] <file|url|->...

`)
		fl.PrintDefaults()
	}
	if err = fl.Parse(args); err != nil {
		return flag.ErrHelp
	}

	args = fl.Args()
	if len(args) == 0 {
		args = []string{flagext.StdIO}
	}

	enc := json.NewEncoder(dst)
	enc.SetIndent(*prefix, *indent)
	enc.SetEscapeHTML(*htmlSafe)
	defer errors.Defer(&err, dst.Close)

	var errs errors.Slice
	for _, arg := range args {
		errs.Push(tidyPrint(arg, enc))
	}
	return errs.Merge()
}

func tidyPrint(arg string, enc *json.Encoder) (err error) {
	src := flagext.FileOrURL(flagext.StdIO, nil)
	if err = src.Set(arg); err != nil {
		return fmt.Errorf("problem with %q: %v\n", arg, err)
	}
	defer errors.Defer(&err, src.Close)

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return fmt.Errorf("problem with %q: %v\n", arg, err)
	}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber() // Preserve number formatting

	var data interface{}

	for dec.More() {
		if err = dec.Decode(&data); err != nil {
			return fmt.Errorf("problem with %q: %v\n", arg, err)
		}

		if err = enc.Encode(&data); err != nil {
			return fmt.Errorf("cannot write out tidy %q: %v\n", arg, err)
		}
	}
	return nil
}
