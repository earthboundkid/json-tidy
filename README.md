# json-tidy
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
Usage of json-tidy:

json-tidy [opts] [file|url|-]
        Gets input (defaults to stdin) and prints clean json to stdout.
  -indent string
        Identation string (default "\t")
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
