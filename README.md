# GoLeRi

A left-right parser for the Go language.

---------------------------------------
  * [Installation](#installation)
  * [Related projects](#related-projects)
  * [Quick usage](#quick-usage)
  
---------------------------------------
## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/transceptor-technology/goleri
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.
## Related projects
---------------------------------------
- [pyleri](https://github.com/transceptor-technology/pyleri): Python parser
- [jsleri](https://github.com/transceptor-technology/jsleri): JavaScript parser
- [cleri](https://github.com/transceptor-technology/cleri): C parser

---------------------------------------
## Quick usage
We recommend using [pyleri](https://github.com/transceptor-technology/pyleri) for creating a grammar and export the grammar to goleri. This way you can create one single grammar and export the grammar to program languages like C, JavaScript, Go and Python.
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

...
// compile your grammar by creating an instance of the Grammar Class.
myGrammar := MyGrammar()

// use the compiled grammar to parse 'strings'
if res, err := myGrammar.Parse("hi 'Iris'"); err == nil {
	fmt.Printf("%v\n", res.IsValid())
	/*
	res.IsValid()
		returns true or false depending if the string is successful parsed
		by the grammar or not.
	res.Pos()  
		returns the position in the string where parsing has end. 
		(if successful this will be equal to the string length)
	res.GetExpecting()
		returns expected elements at position res.Pos(). This can be used
		for auto-completion, auto correction or suggestions.
	res.Tree()
		returns the parse tree.
	*/
}
```
