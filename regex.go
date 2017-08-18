package goleri

import (
	"fmt"
	"regexp"
)

// Regex matches a regular expression.
type Regex struct {
	element
	regex *regexp.Regexp
}

// NewRegex returns a new keyword object.
func NewRegex(gid int, regex *regexp.Regexp) *Regex {
	return &Regex{
		element: element{gid},
		regex:   regex,
	}
}

// GetRegex returns the regular expression
func (regex *Regex) GetRegex() *regexp.Regexp {
	return regex.regex
}

func (regex *Regex) String() string {
	return fmt.Sprintf("<Regex gid:%d regex:%v>", regex.gid, regex.regex)
}

func (regex *Regex) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var nd *Node
	s := p.s[parent.end:]
	m := regex.regex.FindStringIndex(s)
	if m != nil && m[0] == 0 {
		nd = newNode(regex, parent.end)
		nd.end = parent.end + m[1]
		p.appendChild(parent, nd)
	}
	return nd, nil
}
