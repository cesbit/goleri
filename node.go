package goleri

import (
	"fmt"
)

// Node is part or a result tree.
type Node struct {
	Elem     Element
	Start    int
	End      int
	Children []*Node
	Data     interface{} // nil, free to use for anything you like
}

func (n *Node) String() string {
	return fmt.Sprintf("<Node elem:%v children:%d>", n.Elem, len(n.Children))
}

func newNode(elem Element, start int) *Node {
	return &Node{
		Elem:  elem,
		Start: start,
		End:   start,
		Data:  nil,
	}
}
