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

// IsSet returns true if an element has been set.
func (ref *Ref) IsSet() bool { return ref.elem != nil }

func (ref *Ref) String() string {
	return fmt.Sprintf("<Ref isSet:%t>", ref.IsSet())
}

func (ref *Ref) Text() string {
	return fmt.Sprintf("%t", ref.IsSet())
}

// Set reference element.
func (ref *Ref) Set(elem Element) {
	ref.elem = elem
}

func (ref *Ref) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	return ref.elem.parse(p, parent, r)
}
