package goleri

import (
	"fmt"
	"sort"
	"strings"
)

// Tokens matches one of the tokens.
type Tokens struct {
	element
	tokens []string
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// NewTokens returns a new token object.
func NewTokens(gid int, tokens string) *Tokens {
	t := strings.Fields(tokens)
	sort.Sort(sort.Reverse(byLength(t)))
	return &Tokens{
		element: element{gid},
		tokens:  t,
	}
}

// GetToken returns the token
func (tokens *Tokens) GetTokens() []string { return tokens.tokens }

func (tokens *Tokens) String() string {
	return fmt.Sprintf("<Tokens gid:%d tokens:%v>", tokens.gid, tokens.tokens)
}

func (tokens *Tokens) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var nd *Node
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
