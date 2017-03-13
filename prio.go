package goleri

import "fmt"

// PrioMaxRecursionDepth is the maximum recursion depth used.
const PrioMaxRecursionDepth = 10

// Prio can match with a reference to itself.
type Prio struct {
	element
	elements []Element
}

// NewPrio returns a new rule object containing a prio object.
func NewPrio(gid int, elements ...Element) *Rule {
	prio := Prio{
		element:  element{0},
		elements: elements,
	}
	return NewRule(gid, &prio)
}

func (prio *Prio) String() string {
	return fmt.Sprintf("<Prio gid:%d elements:%v>", prio.gid, prio.elements)
}

func (prio *Prio) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {

	if r.depth > PrioMaxRecursionDepth {
		return nil, fmt.Errorf("max recursion depth (%d) is reached", PrioMaxRecursionDepth)
	}

	r.depth++

	for _, elem := range prio.elements {
		nd := newNode(prio, parent.end)
		n, err := p.walk(nd, elem, r, modeRequired)

		if err != nil {
			return nil, err
		}

		if n != nil {
			r.update(n)
		}
	}

	if nd, ok := r.tested[parent.end]; ok && nd != nil {
		p.appendChild(parent, nd)
		return nd, nil
	}

	return nil, nil
}
