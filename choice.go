package goleri

import "fmt"

// Choice must match at least one element.
type Choice struct {
	element
	mostGreedy bool
	elements   []Element
}

// NewChoice returns a new choice object.
func NewChoice(gid int, mostGreedy bool, elements ...Element) *Choice {
	return &Choice{
		element:    element{gid},
		mostGreedy: mostGreedy,
		elements:   elements,
	}
}

// IsMostGreedy return the boolean mostGreedy.
func (choice *Choice) IsMostGreedy() bool { return choice.mostGreedy }

func (choice *Choice) String() string {
	return fmt.Sprintf("<Choice gid:%d elements:%v>", choice.gid, choice.elements)
}

func (choice *Choice) parse(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	if choice.mostGreedy {
		return choice.parseMostGreedy(p, parent, r)
	}
	return choice.parseFirst(p, parent, r)
}

func (choice *Choice) parseMostGreedy(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var mgNode *Node
	var nd *Node

	for _, elem := range choice.elements {
		nd = newNode(choice, parent.End)
		n, err := p.walk(nd, elem, r, modeRequired)
		if err != nil {
			return nil, err
		}
		if n != nil && (mgNode == nil || nd.End > mgNode.End) {
			mgNode = nd
		}
	}

	if mgNode != nil {
		p.appendChild(parent, mgNode)
	}

	return mgNode, nil
}

func (choice *Choice) parseFirst(p *parser, parent *Node, r *ruleStore) (*Node, error) {
	var fNode *Node
	var nd *Node

	for _, elem := range choice.elements {
		nd = newNode(choice, parent.End)
		n, err := p.walk(nd, elem, r, modeRequired)

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
