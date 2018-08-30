# GoLeRi

A left-right parser for the Go language.

---------------------------------------
  * [Installation](#installation)
  * [Related projects](#related-projects)
  * [Quick usage](#quick-usage)
  * [Grammar](#grammar)
    * [Parse](#parse)
	  * [IsValid](#isvalid)
      * [Position](#position)
      * [Tree](#tree)
      * [Expecting](#expecting)
  * [Elements](#elements)
	* [Keyword](#keyword)
	* [Regex](#regex)
	* [Token](#token)
	* [Tokens](#tokens)
	* [Sequence](#sequence)
	* [Choice](#choice)
	* [Repeat](#repeat)
	* [List](#list)
	* [Optional](#optional)
	* [Ref](#ref)
	* [Prio](#prio)


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
- [libcleri](https://github.com/transceptor-technology/libcleri): C parser
- [jleri](https://github.com/transceptor-technology/jleri): Java parser

---------------------------------------
## Quick usage
We recommend using [pyleri](https://github.com/transceptor-technology/pyleri) for creating a grammar and export the grammar to goleri. This way you can create one single grammar and export the grammar to program languages like C, JavaScript, Go and Python.
```go
package main

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
	return goleri.NewGrammar(START, nil)
}

func main() {
	// compile your grammar by creating an instance of the Grammar Class.
	myGrammar := MyGrammar()

	// use the compiled grammar to parse 'strings'
	if res, err := myGrammar.Parse("hi 'Iris'"); err == nil {
		fmt.Printf("%t\n", res.IsValid())
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
}
```


## Grammar

The first Grammar parameter expects a `START` property so the parser knows where to start parsing. The second parameter of Grammar `reKeywords` is set to `nil` in the [Quick Usage](#quickusage) example. In that case goleri uses the default setting of `^\w+` which is equal to `^[A-Za-z0-9_]+`. This default property can be overwritten. (See the [example](#keyword) for keyword). Grammar has a parse method: `Parse()`.

```go
goleri.NewGrammar(elem Element, reKeywords *regexp.Regexp)
```

### Parse
syntax:
```go
MyGrammar.Parse("string")
```
The `parse()` method returns a result object which has the following properties:
- `.GetExpecting()`
- `.isValid()`
- `.Pos()`
- `.Tree()`



#### IsValid

`IsValid()` returns a boolean value, `True` when the given string is valid according to the given grammar, `False` when not valid.

Let us take the example from Quick usage.
```go
if res, err := myGrammar.Parse("hi 'Iris'"); err == nil {
	fmt.Printf("%t\n", res.IsValid())
```

#### Position
`Pos` returns the position where the parser had to stop. (when `IsValid()` is `True` this value will be equal to the length of the given string)

Let us take the example from Quick usage.
```go
if res, err := myGrammar.Parse("hi 'Iris'"); err == nil {
	fmt.Printf("%d\n", res.Pos())
```

#### Tree
`Tree()` contains the parse tree. Even when `IsValid()` is `False` the parse tree is returned but will only contain results as far as parsing has succeeded. The tree is the root node which can include several `children` nodes.

#### Expecting
`GetExpecting()` returns an object containing elements which goleri expects at `Pos`. Even if `IsValid` is true there might be elements in this set, for example when an `Optional` element could be added to the string. Expecting is useful if you want to implement things like auto-completion, syntax error handling, auto-syntax-correction etc.


## Elements
Goleri has several elements which are all subclasses of [Element](#element) and can be used to create a grammar. For every element you can define a Global Identifer (gid) for identifying the element in the parse tree. But the gid is not required and should be set to 0 if not used. Every element has two corresponding methods that can be exported: `Gid()` returns the elements gid and `String()` returns a string including the gid, the element and possible nested elements like "<Keyword gid:0 keyword:hi>".

### Keyword
syntax:
```go
goleri.NewKeyword(gid int, keyword string, ignCase bool)
```
The parser needs to match the keyword which is just a string. When matching keywords we need to tell the parser what characters are allowed in keywords. By default Goleri uses `^\w+` which is equal to `^[A-Za-z0-9_]+`. In this example below we deviate from this default to     `NewKeyword()` accepts a parameter `ignCase` to tell the parser if we should match case insensitive. The following methods return the arguments that are passed to the `NewKeyword()` method: `IsIgnCase()` returns a boolean that indicates if case is ignored and `GetKeyword()` returns the keyword.

Example:

```go
// Let's allow keywords with alphabetic characters and dashes.
START := goleri.NewKeyword(0, "tic-tac-toe", true)
grammar := goleri.NewGrammar(START, regexp.MustCompile(`^[A-Za-z-]+`))
if res, err := grammar.Parse("Tic-Tac-Toe"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", START.GetKeyword()) // tic-tac-toe
	fmt.Printf("%t\n", START.IsIgnCase()) // true
	fmt.Printf("%s\n", START.String()) // <Keyword gid:0 keyword:tic-tac-toe>
}
```

### Regex
syntax:
```go
goleri.NewRegex(gid int, regex *regexp.Regexp)
```
The parser compiles a regular expression using the `regexp` package. The `GetRegex()` method returns the regex that is passed to `NewRegex()`.

Example:

```go
START := goleri.NewRegex(0, regexp.MustCompile(`^(?:'(?:[^']*)')+`))
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("'Iris'"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", START.GetRegex()) // ^(?:'(?:[^']*)')+
	fmt.Printf("%s\n", START.String()) // <Regex gid:0 regex:^(?:'(?:[^']*)')+>
}
```

### Token
syntax:
```go
goleri.NewTokens(gid int, token string)
```
A token can be one or more characters and is usually used to match operators like `+`, `-`, `//` and so on. When we parse a string object where goleri expects an element, it will automatically be converted to a `Token()` object. The `GetToken()` method returns the token that is passed to `NewToken()`.

Example:
```go
START := goleri.NewToken(0, "+")
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("-"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // false
	fmt.Printf("%s\n", START.GetToken()) // +
	fmt.Printf("%s\n", START.String()) // <Token gid:0 token:+>
}
```

### Tokens
syntax:
```go
goleri.NewTokens(gid int, tokens string)
```
Can be used to register multiple tokens at once. The `tokens` argument should be a string with tokens separated by spaces. If given tokens are different in size the parser will try to match the longest tokens first. The `GetTokens()` method returns the tokens that are passed to `NewTokens()`.

Example:
```go
START := goleri.NewTokens(0, "== > < <= >=")
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse(">"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", START.GetTokens()) // [== <= >= > <]
	fmt.Printf("%s\n", START.String()) //<Tokens gid:0 tokens:[== <= >= > <]>
}
```

### Sequence
syntax:
```go
goleri.NewSequence(gid int, elements ...Element)
```
The parser needs to match each element in a sequence.

Example:
```go
START := goleri.NewSequence(
	0,
	goleri.NewKeyword(0, "tic", false),
	goleri.NewKeyword(0, "tac", false),
	goleri.NewKeyword(0, "toe", false))
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("tic tac toe"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", START.String()) /* <Sequence gid:0 elements:[<Keyword gid:0 keyword:tic> <Keyword gid:0 keyword:tac> <Keyword gid:0 keyword:toe>]> */
}
```

### Choice
syntax:
```go
goleri.NewChoice(gid int, mostGreedy bool, elements ...Element)
```
The parser needs to choose between one of the given elements. Choice accepts the parameter `mostGreedy`. When `mostGreedy` is set to `false` the parser will stop at the first match. When `true` the parser will try each element and returns the longest match. Setting `mostGreedy` to `false` can provide some extra performance. Note that the parser will try to match each element in the exact same order they are parsed to Choice. The `IsMostGreedy()` method returns the boolean of mostGreedy.

Example: let us use `Choice` to modify the Quick usage example to allow the string 'bye "Iris"'
```go
choice := goleri.NewChoice(
	0,
	true,
	goleri.NewKeyword(0, "hi", false),
	goleri.NewKeyword(0, "bye", true))
START := goleri.NewSequence(
	0,
	choice,
	goleri.NewRegex(0, regexp.MustCompile(`^(?:'(?:[^']*)')+`)))
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("hi 'Iris'"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
}
if res, err := grammar.Parse("bye 'Iris'"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%t\n", choice.IsMostGreedy()) // true
	fmt.Printf("%s\n", choice.String()) // <Choice gid:0 elements:[<Keyword gid:0 keyword:hi> <Keyword gid:0 keyword:bye>]>
}
```

### Repeat
syntax:
```go
goleri.NewRepeat(gid int, elem Element, min, max int)
```
The parser needs at least `min` elements and at most `max` elements. When `max` is set to 0 we allow unlimited number of elements. `min` can be any integer value equal or higher than 0 but not larger then `max`. The following methods return the min and max arguments that are passed to the `NewRepeat()` method: `GetMin()`, `GetMax()`.

Example:
```go
START := goleri.NewRepeat(
	0,
	goleri.NewKeyword(0, "na", false),
	1,
	8)
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("na na na na na na"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%d\n", START.GetMin()) // 1
	fmt.Printf("%d\n", START.GetMax()) // 8
	fmt.Printf("%s\n", START.String()) // <Repeat gid:0 elem:<Keyword gid:0 keyword:na>>
}
```

### List
syntax:
```go
goleri.NewList(gid int, elem, delimiter Element, min, max int, optClose bool)
```
List is like Repeat but with a delimiter. The delimiter must be defined as an element first, for example as a `NewToken()` or `NewKeyword()`. `min` and `max` have the same meaning as in Repeat. The parameter `optClose` can be set to `true` to allow the list to end with a delimiter. The following methods return the min and max arguments that are passed to the `NewRepeat()` method: `GetMin()`, `GetMax()`.

Example:
```go
START := goleri.NewList(
	0,
	goleri.NewKeyword(0, "ni", false),
	goleri.NewToken(0, ","),
	0,
	8,
	true)
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("ni, ni, ni, ni,"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%d\n", START.GetMin()) // 0
	fmt.Printf("%d\n", START.GetMax()) // 8
	fmt.Printf("%t\n", START.IsOptClose()) // true
	fmt.Printf("%s\n", START.String()) // <List gid:0 elem:<Keyword gid:0 keyword:ni> delimiter:<Token gid:0 token:,>>
}
```

### Optional
syntax:
```go
goleri.NewOptional(gid int, elem Element)
```
The parser looks for an optional element. It is like using `NewRepeat(0, element, 0, 1)` but we encourage to use `Optional` since it is more readable. (and slightly faster)

Example:
```go
optional := goleri.NewOptional(
	0,
	goleri.NewRegex(0, regexp.MustCompile(`^(?:'(?:[^']*)')+`)))
START := goleri.NewSequence(0, goleri.NewKeyword(0, "hi", false), optional)
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("hi"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
}
if res, err := grammar.Parse("hi 'iris'"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", optional.String()) // <Optional gid:0 elem:<Regex gid:0 regex:^(?:'(?:[^']*)')+>>
}
```

### Ref
syntax:
```go
goleri.NewRef()
```
The grammar can make a forward reference to make recursion possible. In the example below we create a forward reference to START but note that a reference to any element can be made.

>Warning: A reference is not protected against testing the same position in
>a string. This could potentially lead to an infinite loop.
>For example:
>```go
>r = goleri.NewRef()
>r = goleri.NewOptional(0, r)  // DON'T DO THIS
>```
>Use [Prio](#prio) if such recursive construction is required.

Example:
```go
START := goleri.NewRef()
START.Set(goleri.NewSequence(
	0,
	goleri.NewToken(0, "["),
	goleri.NewList(
		0,
		goleri.NewChoice(0, true, goleri.NewKeyword(0, "ni", false), START),
		goleri.NewToken(0, ","),
		0,
		0,
		false),
	goleri.NewToken(0, "]")))
grammar := goleri.NewGrammar(START, nil)

if res, err := grammar.Parse("[ni, ni, [ni, [], [ni, ni]]]"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", START.String()) // <Ref isSet:true>
}
```

### Prio
syntax:
```go
goleri.NewPrio(gid int, elements ...Element)
```
Choose the first match from the prio elements and allow `THIS` for recursive operations. With `THIS` we point to the `Prio` element. Probably the example below explains how `Prio` and `THIS` can be used.

>Note: Use a [Ref](#ref) when possible.
>A `Prio` element is required when the same position in a string is potentially
>checked more than once.

Example:
```go
START := goleri.NewPrio(
	0,
	goleri.NewKeyword(0, "ni", false),
	goleri.NewSequence(0, goleri.NewToken(0, "("), goleri.THIS, goleri.NewToken(0, ")")),
	goleri.NewSequence(0, goleri.THIS, goleri.NewKeyword(0, "or", false), goleri.THIS),
	goleri.NewSequence(0, goleri.THIS, goleri.NewKeyword(0, "and", false), goleri.THIS))
grammar := goleri.NewGrammar(START, nil)
if res, err := grammar.Parse("(ni or ni) and (ni or ni)"); err == nil {
	fmt.Printf("%t\n", res.IsValid()) // true
	fmt.Printf("%s\n", START.String()) /* <Prio gid:0 elements:[<Keyword gid:0 keyword:ni> <Sequence gid:0 elements:[<Token gid:0 token:(> <This> <Token gid:0 token:)>]> <Sequence gid:0 elements:[<This> <Keyword gid:0 keyword:or> <This>]> <Sequence gid:0 elements:[<This> <Keyword gid:0 keyword:and> <This>]>]> */
}
```
