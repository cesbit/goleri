package goleri

// Result is used as a Parse return value.
type Result struct {
	isValid    bool
	pos        int
	expect     *expecting
	tree       *Node
	expression string
}

// IsValid returns true when a parse result is valid.
func (r *Result) IsValid() bool { return r.isValid }

// Pos returns the position in the string where parseing has end.
func (r *Result) Pos() int { return r.pos }

// GetExpecting return a list of elements which are expected.
func (r *Result) GetExpecting() []Element {
	return r.expect.getExpecting()
}

// Tree returns the node tree.
func (r *Result) Tree() *Node {
	r.assignExpressionParts()
	return r.tree
}

func (r *Result) assignExpressionParts() {
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		n.Data = r.expression[n.start:n.end]
		for _, c := range n.children {
			walk(c)
		}
	}
	walk(r.tree)
}
