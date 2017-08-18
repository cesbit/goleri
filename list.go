package goleri

import "fmt"

// List must match at least min and at most max times the element and delimiter.
type List struct {
	element
	elem      Element
	delimiter Element
	min       int
	max       int
	optClose  bool // when true the list may end with a delimiter
}

// NewList returns a new list object.
func NewList(gid int, elem, delimiter Element, min, max int, optClose bool) *List {
	return &List{
		element:   element{gid},
		elem:      elem,
		delimiter: delimiter,
		min:       min,
		max:       max,
		optClose:  optClose,
	}
}

func (list *List) String() string {
	return fmt.Sprintf("<List gid:%d elem:%v delimiter:%v>", list.gid, list.elem, list.delimiter)
}

func (list *List) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {

	nd := newNode(list, parent.end)
	i := 0
	j := 0
	for {
		n, err := p.walk(nd, list.elem, r, getMode(i < list.min))
		if err != nil {
			return nil, err
		}
		if n == nil {
			break
		}
		i++
		n, err = p.walk(nd, list.delimiter, r, getMode(i < list.min))
		if err != nil {
			return nil, err
		}
		if n == nil {
			break
		}
		j++
	}

	if i < list.min ||
		(list.max != 0 && i > list.max) ||
		((!list.optClose) && i != 0 && i == j) {
		nd = nil // invalid, make sure nd is nil
	} else {
		p.appendChild(parent, nd)
	}

	return nd, nil
}

func getMode(required bool) uint8 {
	if required {
		return modeRequired
	}
	return modeOptional
}
