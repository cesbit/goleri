package goleri

import (
	"fmt"
	"strings"
)

// Keyword matches a keyword.
type Keyword struct {
	element
	keyword string
	ignCase bool
}

// NewKeyword returns a new keyword object.
func NewKeyword(gid int, keyword string, ignCase bool) *Keyword {
	return &Keyword{
		element: element{gid},
		keyword: keyword,
		ignCase: ignCase,
	}
}

// GetKeyword returns the keyword
func (keyword *Keyword) GetKeyword() string { return keyword.keyword }

// IsIgnCase returns a boolean ignore case
func (keyword *Keyword) IsIgnCase() bool { return keyword.ignCase }

func (keyword *Keyword) String() string {
	return fmt.Sprintf("<Keyword gid:%d keyword:%v>", keyword.gid, keyword.keyword)
}

func (keyword *Keyword) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var match bool
	var nd *Node
	s := p.getKeyword(parent.End)

	if keyword.ignCase {
		match = strings.EqualFold(s, keyword.keyword)
	} else {
		match = strings.Compare(s, keyword.keyword) == 0
	}

	if match {
		nd = newNode(keyword, parent.End)
		nd.End = parent.End + len(keyword.keyword)
		p.appendChild(parent, nd)
	} else {
		p.expect.update(keyword, parent.End)
	}
	return nd, nil
}
