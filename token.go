package goleri

import (
	"fmt"
)

// Token matches a token.
type Token struct {
	element
	token string
	l     int
}

// NewToken returns a new token object.
func NewToken(gid int, token string) *Token {
	return &Token{
		element: element{gid},
		token:   token,
		l:       len(token),
	}
}

// GetToken returns the token
func (token *Token) GetToken() string { return token.token }

func (token *Token) String() string {
	return fmt.Sprintf("<Token gid:%d token:%v>", token.gid, token.token)
}

func (token *Token) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var nd *Node
	match := true
	for i, j := 0, parent.end; i < token.l; i++ {
		if j == p.l || p.s[j] != token.token[i] {
			match = false
			break
		}
		j++
	}

	if match {
		nd = newNode(token, parent.end)
		nd.end = parent.end + token.l
		p.appendChild(parent, nd)
	} else {
		p.expect.update(token, parent.end)
	}

	return nd, nil
}
