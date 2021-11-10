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

func (n *Node) String() string {
	return fmt.Sprintf("<Node elem:%v children:%d>", n.elem, len(n.children))
}
func (n *Node) GetElement() Element {
	return n.elem
}
func (n *Node) GetStart() int {
	return n.start
}
func (n *Node) GetEnd() int {
	return n.end
}
func (n *Node) GetChildren() []*Node {
	return n.children
}

func newNode(elem Element, start int) *Node {
	return &Node{
		elem:  elem,
		start: start,
		end:   start,
		Data:  nil,
	}
}
