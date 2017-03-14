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
        "regexp"

        "github.com/transceptor-technology/goleri"
)

// Element identifiers
const (
        NoGid = iota
        GidKHi = iota
        GidRName = iota
        GidSTART = iota
)

// MyGrammar returns a compiled goleri grammar.
func MyGrammar() *goleri.Grammar {
        rName := goleri.NewRegex(GidRName, regexp.MustCompile(`^(?:'(?:[^']*)')+`))
        kHi := goleri.NewKeyword(GidKHi, "hi", false)
        START := goleri.NewSequence(
                GidSTART,
                kHi,
                rName,
        )
        return goleri.NewGrammar(START, regexp.MustCompile(`^\w+`))
}

// compile your grammar by creating an instance of the Grammar Class.
myGrammar := MyGrammar()

// use the compiled grammar to parse 'strings'
my_grammar.Parse("hi 'Iris'").isValid() // true
my_grammar.Parse("bye 'Iris'").isValid() // false
```
