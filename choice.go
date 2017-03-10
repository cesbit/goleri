package goleri

import "fmt"

// Choice must match at least one element.
type Choice struct {
	element
	mostGreedy bool
	elements   []Element
}

// NewChoice returns a new keyword object.
func NewChoice(gid int, mostGreedy bool, elements ...Element) *Choice {
	return &Choice{
		element:    element{gid},
		mostGreedy: mostGreedy,
		elements:   elements,
	}
}

func (choice *Choice) String() string {
	return fmt.Sprintf("<Choice gid:%d elements:%v>", choice.gid, choice.elements)
}

func (choice *Choice) parse(p *parser, parent *node) (*node, error) {
	if choice.mostGreedy {
		return choice.parseMostGreedy(p, parent)
	}
	return choice.parseMostGreedy(p, parent)
}

func (choice *Choice) parseMostGreedy(p *parser, parent *node) (*node, error) {
	var mgNode *node
	var nd *node

	for _, elem := range choice.elements {
		nd = newNode(choice, parent.end)
		n, err := p.walk(nd, elem, modeRequired)
		if err != nil {
			return nil, err
		}
		if n != nil && (mgNode == nil || nd.end > mgNode.end) {
			mgNode = nd
		}
	}

	if mgNode != nil {
		p.appendChild(parent, mgNode)
	}

	return mgNode, nil
}

func (choice *Choice) parseFirst(p *parser, parent *node) (*node, error) {
	var fNode *node
	var nd *node

	for _, elem := range choice.elements {
		nd = newNode(choice, parent.end)
		n, err := p.walk(nd, elem, modeRequired)

		if err != nil {
			return nil, err
		}

		if n != nil {
			p.appendChild(parent, nd)
			fNode = nd
			break
		}
	}

	return fNode, nil
}
