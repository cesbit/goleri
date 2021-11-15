package goleri

import (
	"fmt"
)

// Node is part or a result tree.
type Node struct {
	elem     Element
	start    int
	end      int
	children []*Node
	Data     interface{} // nil, free to use for anything you like
}

// String returns a representation of the Node
func (n *Node) String() string {
	return fmt.Sprintf("<Node elem:%v children:%d>", n.elem, len(n.children))
}

// GetElement returns the Node element
func (n *Node) GetElement() Element { return n.elem }

// GetStart returns the position in the string where this Node starts
func (n *Node) GetStart() int { return n.start }

// GetEnd returns the position in the string where this Node ends
func (n *Node) GetEnd() int { return n.end }

// GetChildren returns the children of this None
func (n *Node) GetChildren() []*Node { return n.children }

func newNode(elem Element, start int) *Node {
	return &Node{
		elem:  elem,
		start: start,
		end:   start,
		Data:  nil,
	}
}
