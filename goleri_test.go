package goleri

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGoleri(t *testing.T) {
	c := NewRef()
	seq := NewSequence(
		1,
		c,
		NewOptional(0, NewKeyword(0, "iris", false)),
		NewList(0, NewKeyword(0, "ni", true), NewTokens(0, ", - ."), 1, 0, true))

	c.Set(NewChoice(0, true, NewKeyword(0, "hoi", false), NewKeyword(0, "hallo", false)))

	reKw := regexp.MustCompile(`^\w+`)

	gr := NewGrammar(seq, reKw)

	pr, err := gr.Parse("hoi iris ni, ni.ni,")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("IsValid: %v\n", pr.IsValid())
		fmt.Printf("Position: %v\n", pr.Pos())
		fmt.Printf("Expecting: %v\n", pr.GetExpecting())
	}

}
