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

func newNode(elem Element, start int) *Node {
	return &Node{
		elem:  elem,
		start: start,
		end:   start,
		Data:  nil,
	}
}
