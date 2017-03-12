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

func (p *parser) walk(parent *node, elem Element, r *ruleStore, mode uint8) (*node, error) {
	for parent.end < p.l && unicode.IsSpace(rune(p.s[parent.end])) {
		parent.end++
	}

	/* set expecting mode */
	p.expect.setMode(parent.start, mode)

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

func (p *parser) appendChild(parent, child *node) {
	if child.end > p.expect.pos {
		p.expect.empty()
	}
	parent.end = child.end
	parent.children = append(parent.children, child)
}
