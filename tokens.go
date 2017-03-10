package goleri

import (
	"fmt"
	"strings"
)

// Tokens matches one of the tokens.
type Tokens struct {
	element
	tokens []string
}

// NewTokens returns a new token object.
func NewTokens(gid int, tokens string) *Tokens {
	return &Tokens{
		element: element{gid},
		tokens:  strings.Fields(tokens),
	}
}

func (tokens *Tokens) String() string {
	return fmt.Sprintf("<Tokens gid:%d tokens:%v>", tokens.gid, tokens.tokens)
}

func (tokens *Tokens) parse(p *parser, parent *node) (*node, error) {
	var nd *node
	var match bool
	var tokenLen int
	for _, token := range tokens.tokens {
		match = true
		tokenLen = len(token)
		for i, j := 0, parent.end; i < tokenLen; i++ {
			if j == p.l || p.s[j] != token[i] {
				match = false
				break
			}
			j++
		}
		if match {
			break
		}
	}

	if match {
		nd = newNode(tokens, parent.end)
		nd.end = parent.end + tokenLen
		p.appendChild(parent, nd)
	} else {
		p.expect.update(tokens, parent.end)
	}

	return nd, nil
}
