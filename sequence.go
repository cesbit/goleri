package goleri

import "fmt"

// Sequence must match each element in the sequence.
type Sequence struct {
	element
	elements []Element
}

// NewSequence returns a new keyword object.
func NewSequence(gid int, elements ...Element) *Sequence {
	return &Sequence{
		element:  element{gid},
		elements: elements,
	}
}

func (sequence *Sequence) String() string {
	return fmt.Sprintf("<Sequence gid:%d elements:%v>", sequence.gid, sequence.elements)
}

func (sequence *Sequence) parse(p *parser, parent *node) (*node, error) {
	nd := newNode(sequence, parent.end)

	for _, elem := range sequence.elements {
		n, err := p.walk(nd, elem, modeRequired)
		if n == nil || err != nil {
			return nil, err
		}
	}

	p.appendChild(parent, nd)

	return nd, nil
}
