package goleri

import (
	"fmt"
)

// Ref can be used as a forward reference.
type Ref struct {
	element
	elem Element
}

// NewRef returns a new ref object.
func NewRef() *Ref {
	return &Ref{
		element: element{0},
		elem:    nil,
	}
}

func (ref *Ref) String() string {
	return fmt.Sprintf("<Ref elem:%v>", ref.elem)
}

// Set reference element.
func (ref *Ref) Set(elem Element) {
	ref.elem = elem
}

func (ref *Ref) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	return ref.elem.parse(p, parent, r)
}
