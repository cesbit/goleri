package goleri

// Rule is a prio wrapper..
type Rule struct {
	element
	elem Element
}

type ruleStore struct {
	tested map[int]*Node
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
	return rule.elem.String()
}

func (rule *Rule) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {

	nd := newNode(rule, parent.end)

	rs := ruleStore{make(map[int]*Node), rule.elem, 0}

	n, err := p.walk(nd, rs.root, &rs, modeRequired)
	if n == nil || err != nil {
		return nil, err
	}

	p.appendChild(parent, nd)

	return nd, nil
}

func (r *ruleStore) update(nd *Node) {
	if n, ok := r.tested[nd.start]; ok {
		if n == nil || nd.end > n.end {
			r.tested[nd.start] = nd
			return
		}
	} else {
		r.tested[nd.start] = nd
	}
}
