package goleri

import "fmt"

// Optional can match one element.
type Optional struct {
	element
	elem Element
}

// NewOptional returns a new optional object.
func NewOptional(gid int, elem Element) *Optional {
	return &Optional{
		element: element{gid},
		elem:    elem,
	}
}

func (optional *Optional) String() string {
	return fmt.Sprintf("<Optional gid:%d elem:%v>", optional.gid, optional.elem)
}

func (optional *Optional) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	nd := newNode(optional, parent.end)
	n, err := p.walk(nd, optional.elem, r, modeOptional)

	if err != nil {
		return nil, err
	}

	if n != nil {
		p.appendChild(parent, nd)
	}

	return nd, nil
}
