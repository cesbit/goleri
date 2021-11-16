package goleri

// Element interface for all goleri elements.
type Element interface {
	Gid() int
	String() string
	parse(p *parser, parent *Node, r *ruleStore) (*Node, error)
	Text() string
}

type element struct {
	gid int
}

func (elem *element) Gid() int {
	return elem.gid
}
