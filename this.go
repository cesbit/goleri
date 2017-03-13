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

func (t *This) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var nd *Node
	var ok bool
	if nd, ok = r.tested[parent.end]; !ok {
		nd = newNode(t, parent.end)

		r.tested[parent.end] = nil

		n, err := p.walk(nd, r.root, r, modeRequired)
		if n == nil || err != nil {
			return nil, err
		}
		r.tested[parent.end] = n
	}

	if nd != nil {
		p.appendChild(parent, nd)
	}

	return nd, nil
}
