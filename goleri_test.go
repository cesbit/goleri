package goleri

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	t.Helper()

	if reflect.TypeOf(a).Kind() == reflect.Slice {
		va := reflect.ValueOf(a)
		vb := reflect.ValueOf(b)
		if va.Len() == 0 && vb.Len() == 0 {
			return
		}
	}

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

func TestKeywordIgnCase(t *testing.T) {
	hi := NewKeyword(0, "hi", true)
	grammar := NewGrammar(hi, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, hi.Gid())
	assertEquals(t, true, hi.IsIgnCase())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, " hi ").IsValid())
	assertEquals(t, true, parse(t, grammar, "Hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hello").IsValid())
	assertEquals(t, "<Keyword gid:0 keyword:hi>", hi.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestSequence(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	iris := NewKeyword(0, "iris", false)
	sequence := NewSequence(0, hi, iris)
	grammar := NewGrammar(sequence, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, sequence.Gid())
	assertEquals(t, true, parse(t, grammar, "hi iris").IsValid())
	assertEquals(t, true, parse(t, grammar, " hi iris ").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi siri").IsValid())
	assertEquals(t, false, parse(t, grammar, "hello iris").IsValid())
	assertEquals(t, "<Sequence gid:0 elements:[<Keyword gid:0 keyword:hi> <Keyword gid:0 keyword:iris>]>", sequence.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi iris").GetExpecting())
	assertEquals(t, []Element{iris}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestChoiceMostGreedy(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	iris := NewKeyword(0, "iris", false)
	sequence := NewSequence(0, hi, iris)
	choice := NewChoice(0, true, hi, sequence)
	grammar := NewGrammar(choice, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, choice.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi iris").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi siri").IsValid())
	assertEquals(t, "<Choice gid:0 elements:[<Keyword gid:0 keyword:hi> <Sequence gid:0 elements:[<Keyword gid:0 keyword:hi> <Keyword gid:0 keyword:iris>]>]>", choice.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi iris").GetExpecting())
	assertEquals(t, []Element{iris}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestChoiceFirstMatch(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	iris := NewKeyword(0, "iris", false)
	sequence := NewSequence(0, hi, iris)
	choice := NewChoice(0, false, hi, sequence)
	grammar := NewGrammar(choice, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, choice.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi iris").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi siri").IsValid())
	assertEquals(t, "<Choice gid:0 elements:[<Keyword gid:0 keyword:hi> <Sequence gid:0 elements:[<Keyword gid:0 keyword:hi> <Keyword gid:0 keyword:iris>]>]>", choice.String())
	assertEquals(t, []Element{EOS}, parse(t, grammar, "hi iris").GetExpecting())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestOptional(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	optional := NewOptional(0, hi)
	grammar := NewGrammar(optional, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, optional.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "hello").IsValid())
	assertEquals(t, "<Optional gid:0 elem:<Keyword gid:0 keyword:hi>>", optional.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestToken(t *testing.T) {
	token := NewToken(0, "+")
	grammar := NewGrammar(token, regexp.MustCompile(`^\W+`)) //??

	// assert statements
	assertEquals(t, 0, token.Gid())
	assertEquals(t, true, parse(t, grammar, "+").IsValid())
	assertEquals(t, true, parse(t, grammar, " + ").IsValid())
	assertEquals(t, false, parse(t, grammar, "++").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Token gid:0 token:+>", token.String())
	assertEquals(t, []Element{}, parse(t, grammar, "+").GetExpecting())
	assertEquals(t, []Element{token}, parse(t, grammar, "").GetExpecting())
}

func TestTokenMultiChars(t *testing.T) {
	token := NewToken(0, "+=")
	grammar := NewGrammar(token, regexp.MustCompile(`^\W+`)) //??

	// assert statements
	assertEquals(t, 0, token.Gid())
	assertEquals(t, true, parse(t, grammar, "+=").IsValid())
	assertEquals(t, true, parse(t, grammar, " += ").IsValid())
	assertEquals(t, false, parse(t, grammar, "+").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Token gid:0 token:+=>", token.String())
	assertEquals(t, []Element{}, parse(t, grammar, "+=").GetExpecting())
	assertEquals(t, []Element{token}, parse(t, grammar, "").GetExpecting())
}

func TestList(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	token := NewToken(0, ",")
	list := NewList(0, hi, token, 1, 3, false)
	grammar := NewGrammar(list, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, list.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi, hi, hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi , hi , hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi , hi , hi,").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi, hi, hi, hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi hi").IsValid())
	assertEquals(t, false, parse(t, grammar, ", ").IsValid())
	assertEquals(t, "<List gid:0 elem:<Keyword gid:0 keyword:hi> delimiter:<Token gid:0 token:,>>", list.String())
	assertEquals(t, []Element{token}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestListEndDelimiter(t *testing.T) {
	hi := NewKeyword(0, "hi", true)
	token := NewToken(0, ",")
	list := NewList(0, hi, token, 1, 3, true)
	grammar := NewGrammar(list, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, list.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi, hi, hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi , hi , hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi , hi , hi,").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi, hi, hi, hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi hi").IsValid())
	assertEquals(t, false, parse(t, grammar, ", ").IsValid())
	assertEquals(t, "<List gid:0 elem:<Keyword gid:0 keyword:hi> delimiter:<Token gid:0 token:,>>", list.String())
	assertEquals(t, []Element{token}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestRepeat(t *testing.T) {
	hi := NewKeyword(0, "hi", true)
	repeat := NewRepeat(0, hi, 1, 3)
	grammar := NewGrammar(repeat, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, repeat.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi hi hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi  hi  hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi hi hi hi").IsValid())
	assertEquals(t, "<Repeat gid:0 elem:<Keyword gid:0 keyword:hi>>", repeat.String())
	assertEquals(t, []Element{hi}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestTokens(t *testing.T) {
	tokens := NewTokens(0, "== != >=   >   < <=")
	grammar := NewGrammar(tokens, regexp.MustCompile(`^\W+`)) //??

	// assert statements
	assertEquals(t, 0, tokens.Gid())
	assertEquals(t, true, parse(t, grammar, "==").IsValid())
	assertEquals(t, true, parse(t, grammar, "<=").IsValid())
	assertEquals(t, true, parse(t, grammar, ">").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Tokens gid:0 tokens:[== != >= <= > <]>", tokens.String())
	assertEquals(t, []Element{}, parse(t, grammar, "==").GetExpecting())
	assertEquals(t, []Element{tokens}, parse(t, grammar, "").GetExpecting())
}

func TestRegex(t *testing.T) {
	regex := NewRegex(0, regexp.MustCompile("^(/[^/\\\\]*(?:\\\\.[^/\\\\]*)*/i?)"))
	grammar := NewGrammar(regex, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, 0, regex.Gid())
	assertEquals(t, true, parse(t, grammar, "/hi/").IsValid())
	assertEquals(t, true, parse(t, grammar, "/hi/i").IsValid())
	assertEquals(t, true, parse(t, grammar, " //i").IsValid())
	assertEquals(t, false, parse(t, grammar, "x//i").IsValid())
	assertEquals(t, false, parse(t, grammar, "/hi//hi/i").IsValid())
	assertEquals(t, false, parse(t, grammar, "//x").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Regex gid:0 regex:^(/[^/\\\\]*(?:\\\\.[^/\\\\]*)*/i?)>", regex.String())
	assertEquals(t, []Element{}, parse(t, grammar, "/hi/i").GetExpecting())
	assertEquals(t, []Element{}, parse(t, grammar, "").GetExpecting())
}

func TestRef(t *testing.T) {
	ref := NewRef()
	hi := NewKeyword(0, "hi", false)
	grammar := NewGrammar(ref, regexp.MustCompile(`^\w+`))
	ref.Set(hi)

	// assert statements
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Ref elem:<Keyword gid:0 keyword:hi>>", ref.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
}

func TestPrio(t *testing.T) {
	prio := NewPrio(1,
		NewKeyword(0, "hi", false),
		NewKeyword(0, "bye", false),
		NewSequence(0, NewToken(0, "("), THIS, NewToken(0, ")")),
		NewSequence(0, THIS, NewKeyword(0, "or", false), THIS),
		NewSequence(0, THIS, NewKeyword(0, "and", false), THIS))

	grammar := NewGrammar(prio, regexp.MustCompile(`^\w+`))

	// assert statements
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "(bye)").IsValid())
	assertEquals(t, true, parse(t, grammar, "(hi and bye)").IsValid())
	assertEquals(t, true, parse(t, grammar, "(hi or hi) and (hi or hi)").IsValid())
	assertEquals(t, true, parse(t, grammar, "(hi or (hi and bye))").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "(hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "()").IsValid())
	assertEquals(t, false, parse(t, grammar, "(hi or hi) and").IsValid())

	_, err := grammar.Parse("(((((((((((((((((((((((((((((((((((((((((((((((((((hi)))))))))))))))))))))))))))))))))))))))))))))))))))")
	assertEquals(t, fmt.Errorf("max recursion depth (50) is reached"), err)
	assertEquals(t, "<Prio gid:1 elements:[<Keyword gid:0 keyword:hi> <Keyword gid:0 keyword:bye> <Sequence gid:0 elements:[<Token gid:0 token:(> <This> <Token gid:0 token:)>]> <Sequence gid:0 elements:[<This> <Keyword gid:0 keyword:or> <This>]> <Sequence gid:0 elements:[<This> <Keyword gid:0 keyword:and> <This>]>]>", prio.String())
}
