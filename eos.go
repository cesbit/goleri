package goleri

// EndOfStatement is used in expected.
type EndOfStatement struct {
	element
}

// EOS is used as End-Of-Statement element.
var EOS = new(EndOfStatement)

func (eos EndOfStatement) String() string {
	return "<EndOfStatement>"
}

func (eos EndOfStatement) Text() string {
	return "<EndOfStatement>"
}

func (eos EndOfStatement) parse(p *parser, n *Node, r *ruleStore) (*Node, error) {
	return nil, nil
}
