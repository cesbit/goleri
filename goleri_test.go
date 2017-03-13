package goleri

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGoleri(t *testing.T) {
	c := NewRef()
	ni := NewRegex(0, regexp.MustCompile(`(?i)^n[ioa]`))

	calc := NewPrio(
		4,
		ni,
		NewSequence(0, THIS, NewTokens(0, "+ - / %"), THIS),
		NewSequence(0, NewToken(0, "("), THIS, NewToken(0, ")")),
	)

	seq := NewSequence(
		1,
		calc,
		c,
		NewOptional(0, NewKeyword(0, "iris", false)),
		NewList(0, ni, NewTokens(0, ", - ."), 1, 0, true))

	c.Set(NewChoice(0, true, NewKeyword(0, "hoi", false), NewKeyword(0, "hallo", false)))

	reKw := regexp.MustCompile(`^\w+`)

	gr := NewGrammar(seq, reKw)

	pr, err := gr.Parse("(Ni+(no - na)) hoi iris ni, ni.ni,")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("IsValid: %v\n", pr.IsValid())
		fmt.Printf("Position: %v\n", pr.Pos())
		fmt.Printf("Expecting: %v\n", pr.GetExpecting())
		fmt.Printf("Tree: %v\n", pr.Tree())
		if pr.IsValid() {
			fmt.Printf("Tree: %v\n", pr.Tree().children[0])
		}
	}

}
