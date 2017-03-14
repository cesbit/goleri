# GoLeRi

A left-right parser for the Go language.

---------------------------------------
  * [Installation](#installation)
  * [Quick usage](#quick usage)
  
---------------------------------------

## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/transceptor-technology/goleri
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Quick usage
```go
package grammar

import (
	"fmt"
	"regexp"

	"github.com/transceptor-technology/goleri"
)

// Element identifiers
const (
	NoGid    = iota
	GidHi    = iota
	GidName  = iota
	GidSTART = iota
)

// MyGrammar returns a compiled goleri grammar.
func MyGrammar() *goleri.Grammar {
	name := goleri.NewRegex(GidName, regexp.MustCompile(`^(?:'(?:[^']*)')+`))
	hi := goleri.NewKeyword(GidHi, "hi", false)
	START := goleri.NewSequence(GidSTART, hi, name)
	return goleri.NewGrammar(START, regexp.MustCompile(`^\w+`))
}

func main() {
	// compile your grammar by creating an instance of the Grammar Class.
	myGrammar := MyGrammar()

	// use the compiled grammar to parse 'strings'
	if res, err := myGrammar.Parse("hi 'Iris'"); err == nil {
		fmt.Printf("%v\n", res.IsValid())
	}
}
```
