package goleri

// Element interface for all goleri elements.
type Element interface {
	Gid() int
	parse(p *parser, n *node) (*node, error)
}

type element struct {
	gid int
}

func (elem *element) Gid() int {
	return elem.gid
}
