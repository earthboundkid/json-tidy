# json-tidy [![GoDoc](https://godoc.org/github.com/carlmjohnson/json-tidy?status.svg)](https://godoc.org/github.com/carlmjohnson/json-tidy) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/json-tidy)](https://goreportcard.com/report/github.com/carlmjohnson/json-tidy)
Pretty prints JSON

## Installation
First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```shell
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/carlmjohnson/json-tidy
```

## Screenshots
```shell
$ json-tidy -h
Gets input files and URLs (defaults to stdin) and outputs tidy JSON.

Usage of json-tidy:

json-tidy [opts] <file|url|->...
  -html-safe
        Escape special characters for easy embedding in HTML
  -indent string
        Identation string (default "\t")
  -output file
        write tidy JSON to file (default stdout)
  -prefix string
        Prefix string
```

```shell
$ echo '{"a": 1, "b": [true, false]}' | json-tidy
{
        "a": 1,
        "b": [
                true,
                false
        ]
}
```
