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
	grammar := NewGrammar(hi, nil)

	// assert statements
	assertEquals(t, 0, hi.Gid())
	assertEquals(t, "hi", hi.GetKeyword())
	assertEquals(t, false, hi.IsIgnCase())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, " hi ").IsValid())
	assertEquals(t, false, parse(t, grammar, "Hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "hello").IsValid())
	assertEquals(t, "<Keyword gid:0 keyword:hi>", hi.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 2, parse(t, grammar, "hi").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi").Tree().String())
}

func TestKeywordIgnCase(t *testing.T) {
	hi := NewKeyword(0, "hi", true)
	grammar := NewGrammar(hi, nil)

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
	grammar := NewGrammar(sequence, nil)

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
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 7, parse(t, grammar, "hi iris").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi iris").Tree().String())
}

func TestChoiceMostGreedy(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	iris := NewKeyword(0, "iris", false)
	sequence := NewSequence(0, hi, iris)
	choice := NewChoice(0, true, hi, sequence)
	grammar := NewGrammar(choice, nil)

	// assert statements
	assertEquals(t, true, choice.IsMostGreedy())
	assertEquals(t, 0, choice.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi iris").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi siri").IsValid())
	assertEquals(t, "<Choice gid:0 elements:[<Keyword gid:0 keyword:hi> <Sequence gid:0 elements:[<Keyword gid:0 keyword:hi> <Keyword gid:0 keyword:iris>]>]>", choice.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi iris").GetExpecting())
	assertEquals(t, []Element{iris}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 2, parse(t, grammar, "hi").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi").Tree().String())
}

func TestChoiceFirstMatch(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	iris := NewKeyword(0, "iris", false)
	sequence := NewSequence(0, hi, iris)
	choice := NewChoice(0, false, hi, sequence)
	grammar := NewGrammar(choice, nil)

	// assert statements
	assertEquals(t, 0, choice.Gid())
	assertEquals(t, false, choice.IsMostGreedy())
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
	grammar := NewGrammar(optional, nil)

	// assert statements
	assertEquals(t, 0, optional.Gid())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "hello").IsValid())
	assertEquals(t, "<Optional gid:0 elem:<Keyword gid:0 keyword:hi>>", optional.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 2, parse(t, grammar, "hi").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi").Tree().String())
}

func TestToken(t *testing.T) {
	token := NewToken(0, "+")
	grammar := NewGrammar(token, nil)

	// assert statements
	assertEquals(t, 0, token.Gid())
	assertEquals(t, "+", token.GetToken())
	assertEquals(t, true, parse(t, grammar, "+").IsValid())
	assertEquals(t, true, parse(t, grammar, " + ").IsValid())
	assertEquals(t, false, parse(t, grammar, "++").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Token gid:0 token:+>", token.String())
	assertEquals(t, []Element{}, parse(t, grammar, "+").GetExpecting())
	assertEquals(t, []Element{token}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 1, parse(t, grammar, "+").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "+").Tree().String())
}

func TestTokenMultiChars(t *testing.T) {
	token := NewToken(0, "+=")
	grammar := NewGrammar(token, nil)

	// assert statements
	assertEquals(t, 0, token.Gid())
	assertEquals(t, true, parse(t, grammar, "+=").IsValid())
	assertEquals(t, true, parse(t, grammar, " += ").IsValid())
	assertEquals(t, false, parse(t, grammar, "+").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Token gid:0 token:+=>", token.String())
	assertEquals(t, []Element{}, parse(t, grammar, "+=").GetExpecting())
	assertEquals(t, []Element{token}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 2, parse(t, grammar, "+=").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "+=").Tree().String())
}

func TestTokens(t *testing.T) {
	tokens := NewTokens(0, "== != >=   >   < <=")
	grammar := NewGrammar(tokens, nil)

	// assert statements
	assertEquals(t, 0, tokens.Gid())
	assertEquals(t, []string{"==", "!=", ">=", "<=", ">", "<"}, tokens.GetTokens()) // ??
	assertEquals(t, true, parse(t, grammar, "==").IsValid())
	assertEquals(t, true, parse(t, grammar, "<=").IsValid())
	assertEquals(t, true, parse(t, grammar, ">").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, "<Tokens gid:0 tokens:[== != >= <= > <]>", tokens.String())
	assertEquals(t, []Element{}, parse(t, grammar, "==").GetExpecting())
	assertEquals(t, []Element{tokens}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 2, parse(t, grammar, "==").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "==").Tree().String())
}

func TestList(t *testing.T) {
	hi := NewKeyword(0, "hi", false)
	token := NewToken(0, ",")
	list := NewList(0, hi, token, 1, 3, false)
	grammar := NewGrammar(list, nil)

	// assert statements
	assertEquals(t, 0, list.Gid())
	assertEquals(t, false, list.IsOptClose())
	assertEquals(t, 1, list.GetMin())
	assertEquals(t, 3, list.GetMax())
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
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 10, parse(t, grammar, "hi, hi, hi").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi, hi, hi").Tree().String())
}

func TestListEndDelimiter(t *testing.T) {
	hi := NewKeyword(0, "hi", true)
	token := NewToken(0, ",")
	list := NewList(0, hi, token, 1, 3, true)
	grammar := NewGrammar(list, nil)

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

func TestChildren(t *testing.T) {
	grammar := NewGrammar(NewSequence(0, NewKeyword(0, "hi", true), NewKeyword(0, "iris", true)), nil)
	assertEquals(t, true, parse(t, grammar, "hi iris").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi iris").Tree().Data == nil)
	assertEquals(t, "hi iris", parse(t, grammar, "hi iris").Tree().Data.(string))
	// the root node has 1 child containing the sequence
	assertEquals(t, 1, len(parse(t, grammar, "hi iris").Tree().GetChildren()))
	assertEquals(t, false, parse(t, grammar, "hi iris").Tree().GetChildren()[0].Data == nil)
	assertEquals(t, 2, len(parse(t, grammar, "hi iris").Tree().GetChildren()[0].GetChildren()))
	assertEquals(t, "hi", parse(t, grammar, "hi iris").Tree().GetChildren()[0].GetChildren()[0].Data.(string))
	assertEquals(t, false, parse(t, grammar, "hi iris").Tree().GetChildren()[0].GetChildren()[1].Data == nil)
	assertEquals(t, "iris", parse(t, grammar, "hi iris").Tree().GetChildren()[0].GetChildren()[1].Data.(string))
}

func TestRepeat(t *testing.T) {
	hi := NewKeyword(0, "hi", true)
	repeat := NewRepeat(0, hi, 1, 3)
	grammar := NewGrammar(repeat, nil)

	// assert statements
	assertEquals(t, 0, repeat.Gid())
	assertEquals(t, 1, repeat.GetMin())
	assertEquals(t, 3, repeat.GetMax())
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi hi hi").IsValid())
	assertEquals(t, true, parse(t, grammar, "hi  hi  hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, false, parse(t, grammar, "hi hi hi hi").IsValid())
	assertEquals(t, "<Repeat gid:0 elem:<Keyword gid:0 keyword:hi>>", repeat.String())
	assertEquals(t, []Element{hi}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 8, parse(t, grammar, "hi hi hi").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi hi hi").Tree().String())
}

func TestRegex(t *testing.T) {
	regex := NewRegex(0, regexp.MustCompile(`^(/[^/\\]*(?:\\.[^/\\]*)*/i?)`))
	grammar := NewGrammar(regex, nil)

	// assert statements
	assertEquals(t, 0, regex.Gid())
	assertEquals(t, regexp.MustCompile(`^(/[^/\\]*(?:\\.[^/\\]*)*/i?)`), regex.GetRegex()) // ??
	assertEquals(t, true, parse(t, grammar, "/hi/").IsValid())
	assertEquals(t, true, parse(t, grammar, "/hi/i").IsValid())
	assertEquals(t, true, parse(t, grammar, " //i").IsValid())
	assertEquals(t, false, parse(t, grammar, "x//i").IsValid())
	assertEquals(t, false, parse(t, grammar, "/hi//hi/i").IsValid())
	assertEquals(t, false, parse(t, grammar, "//x").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, `<Regex gid:0 regex:^(/[^/\\]*(?:\\.[^/\\]*)*/i?)>`, regex.String())
	assertEquals(t, []Element{}, parse(t, grammar, "/hi/i").GetExpecting())
	assertEquals(t, []Element{regex}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 5, parse(t, grammar, "/hi/i").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "/hi/i").Tree().String())

}

func TestRef(t *testing.T) {
	ref := NewRef()
	hi := NewKeyword(0, "hi", false)
	grammar := NewGrammar(ref, nil)

	// assert statements (before set)
	assertEquals(t, false, ref.IsSet())
	assertEquals(t, "<Ref isSet:false>", ref.String())

	ref.Set(hi)

	// assert statements
	assertEquals(t, true, parse(t, grammar, "hi").IsValid())
	assertEquals(t, false, parse(t, grammar, "").IsValid())
	assertEquals(t, true, ref.IsSet())
	assertEquals(t, "<Ref isSet:true>", ref.String())
	assertEquals(t, []Element{}, parse(t, grammar, "hi").GetExpecting())
	assertEquals(t, []Element{hi}, parse(t, grammar, "").GetExpecting())
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 2, parse(t, grammar, "hi").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "hi").Tree().String())

}

func TestPrio(t *testing.T) {
	prio := NewPrio(1,
		NewKeyword(0, "hi", false),
		NewKeyword(0, "bye", false),
		NewSequence(0, NewToken(0, "("), THIS, NewToken(0, ")")),
		NewSequence(0, THIS, NewKeyword(0, "or", false), THIS),
		NewSequence(0, THIS, NewKeyword(0, "and", false), THIS))

	grammar := NewGrammar(prio, nil)

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
	assertEquals(t, 0, parse(t, grammar, "").Pos())
	assertEquals(t, 25, parse(t, grammar, "(hi or hi) and (hi or hi)").Pos())
	assertEquals(t, "<Node elem:<nil> children:1>", parse(t, grammar, "(hi or hi) and (hi or hi)").Tree().String())
}
