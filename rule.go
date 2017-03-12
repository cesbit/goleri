package goleri

import "fmt"

// Rule is a prio wrapper..
type Rule struct {
	element
	elem Element
}

type ruleStore struct {
	tested []*node
	root   Element
	depth  int
}

// NewRule returns a new keyword object.
func NewRule(gid int, elem Element) *Rule {
	return &Rule{
		element: element{gid},
		elem:    elem,
	}
}

func (rule *Rule) String() string {
	return fmt.Sprintf("<Rule gid:%d elem:%v>", rule.gid, rule.elem)
}

func (rule *Rule) parse(p *parser, parent *node, r *ruleStore) (*node, error) {

	nd := newNode(rule, parent.end)

	rs := ruleStore{[]*node{}, rule.elem, 0}

	n, err := p.walk(nd, rs.root, &rs, modeRequired)
	if n == nil || err != nil {
		return nil, err
	}

	p.appendChild(parent, nd)

	return nd, nil
}
