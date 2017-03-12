package goleri

// This is used in a prio.
type This struct {
	element
}

// THIS is used as This element.
var THIS = new(This)

func (t This) String() string {
	return "<This>"
}

func (t This) parse(p *parser, n *node, r *ruleStore) (*node, error) {
	return nil, nil
}
