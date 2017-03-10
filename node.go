package goleri

type node struct {
	elem     Element
	start    int
	end      int
	children []*node
	Data     interface{} // nil, free to use for anything you like
}

func newNode(elem Element, start int) *node {
	return &node{
		elem:  elem,
		start: start,
		end:   start,
		Data:  nil,
	}
}
