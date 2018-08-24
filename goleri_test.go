package goleri

import (
	"reflect"
	"regexp"
	"testing"
)

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%s != %s", a, b)
	}
}

func parse(t *testing.T, g *Grammar, s string) *Result {
	t.Helper()
	res, err := g.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestKeyword(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	grammar := NewGrammar(hi, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, hi.Gid())
	assertEquals(t, false, hi.IsIgnCase())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, " hi ").IsValid())
	assertEquals(t, false, parse(t, grammar, "Hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hello").IsValid())
	assertEquals(t, "<Keyword gid:0 keyword:hi>", hi.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}
