package goleri

import "fmt"

// Repeat must match at least min and at most max times the element.
type Repeat struct {
	element
	elem Element
	min  int
	max  int
}

// NewRepeat returns a new repeat object.
func NewRepeat(gid int, elem Element, min, max int) *Repeat {
	return &Repeat{
		element: element{gid},
		elem:    elem,
		min:     min,
		max:     max,
	}
}

// GetMin returns the min
func (repeat *Repeat) GetMin() int { return repeat.min }

// GetMax returns the max
func (repeat *Repeat) GetMax() int { return repeat.max }

func (repeat *Repeat) String() string {
	return fmt.Sprintf("<Repeat gid:%d elem:%v>", repeat.gid, repeat.elem)
}

func (repeat *Repeat) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {

	nd := newNode(repeat, parent.End)
	var i int
	for i = 0; repeat.max == 0 || i < repeat.max; i++ {
		n, err := p.walk(nd, repeat.elem, r, modeRequired)
		if err != nil {
			return nil, err
		}
		if n == nil {
			break
		}
	}

	if i < repeat.min {
		nd = nil // invalid, make sure nd is nil
	} else {
		p.appendChild(parent, nd)
	}

	return nd, nil
}
