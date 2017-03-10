package goleri

import (
	"fmt"
	"regexp"
	"unicode"
)

// Grammar is the grammar entry point.
type Grammar struct {
	elem       Element
	reKeywords *regexp.Regexp
}

// NewGrammar returns a new grammar type.
func NewGrammar(elem Element, reKeywords *regexp.Regexp) *Grammar {
	return &Grammar{
		elem:       elem,
		reKeywords: reKeywords,
	}
}

// Parse grammar.
func (g *Grammar) Parse(s string) (*Result, error) {
	var pr *Result

	p := newParser(s, g.reKeywords)
	nd := newNode(nil, 0)
	n, err := p.walk(nd, g.elem, modeRequired)
	if err != nil {
		return nil, err
	}

	pr = &Result{n != nil, nd.end, p.expect}
	end := p.l

	// ignore white space at end
	for end > 0 && unicode.IsSpace(rune(p.s[end-1])) {
		end--
	}

	fmt.Printf("end: %d node: %d\n", end, nd.end)

	if nd.end < end {
		pr.isValid = false
	}

	if nd.end < end && len(p.expect.required) == 0 {
		p.expect.setMode(nd.end, modeRequired)
		p.expect.update(EOS, nd.end)
	}

	return pr, nil
}
