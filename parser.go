package goleri

import (
	"regexp"
	"unicode"
)

type parser struct {
	s          string
	l          int
	kwc        map[int]string
	reKeywords *regexp.Regexp
	expect     *expecting
}

func newParser(s string, reKeywords *regexp.Regexp) *parser {
	return &parser{
		s:          s,
		l:          len(s),
		kwc:        make(map[int]string),
		reKeywords: reKeywords,
		expect:     newExpecting(),
	}
}

func (p *parser) walk(parent *Node, elem Element, r *ruleStore, mode uint8) (*Node, error) {
	for parent.End < p.l && unicode.IsSpace(rune(p.s[parent.End])) {
		parent.End++
	}

	/* set expecting mode */
	p.expect.setMode(parent.Start, mode)

	return elem.parse(p, parent, r)
}

func (p *parser) getKeyword(pos int) string {
	if s, ok := p.kwc[pos]; ok {
		return s
	}
	s := p.reKeywords.FindString(p.s[pos:])
	p.kwc[pos] = s
	return s
}

func (p *parser) appendChild(parent, child *Node) {
	if child.End > p.expect.pos {
		p.expect.empty()
	}
	parent.End = child.End
	parent.Children = append(parent.Children, child)
}
