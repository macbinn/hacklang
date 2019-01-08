package token

import "testing"

func TestParser_NextToken(t *testing.T) {
	cases := []struct {
		Text  string
		Type  Type
		Value string
	}{
		{"abc", ID, "abc"},
		{"1234", NUMBER, "1234"},
		{"true", BOOL, "true"},
		{"[", LBRACK, "["},
		{"]", RBRACK, "]"},
		{"{", LBRACE, "{"},
		{"}", RBRACE, "}"},
		{":", COLON, ":"},
		{",", COMMA, ","},
		{"`abc`\n\n ", STRING, "abc"},
		{"`abc`", STRING, "abc"},
	}
	for _, c := range cases {
		p := Parser{
			buf: []byte(c.Text),
			pos: 0,
		}
		token, err := p.NextToken()
		if err != nil {
			t.Fatalf("%s: next token error %s", c.Text, err)
		}
		if token.Type != c.Type {
			t.Errorf("%s: expect %s but got %s", c.Text, c.Type, token.Type)
		}
		if token.Value != c.Value {
			t.Errorf("%s: expect %s but got %s", c.Text, c.Value, token.Value)
		}
	}
}

func TestParser_Parse(t *testing.T) {
	cases := []struct {
		code  string
		types []Type
	}{
		{
			code:  "a = 1",
			types: []Type{ID, EQUAL, NUMBER},
		},
		{
			code:  "a = true",
			types: []Type{ID, EQUAL, BOOL},
		},
		{
			code:  "a = i => { print(i) }",
			types: []Type{ID, EQUAL, ID, EQUAL, RARROW, LBRACE, ID, LPAREN, ID, RPAREN, RBRACE},
		},
		{
			code:  "a.forEach(1)",
			types: []Type{ID, DOT, ID, LPAREN, NUMBER, RPAREN},
		},
		{
			code:  "a: i=>{ print(i) }",
			types: []Type{ID, COLON, ID, EQUAL, RARROW, LBRACE, ID, LPAREN, ID, RPAREN, RBRACE},
		},
		{
			code:  "a.f()()",
			types: []Type{ID, DOT, ID, LPAREN, RPAREN, LPAREN, RPAREN},
		},
		{
			code: "a\n\nb c\n",
			types: []Type{ID, ID, ID},
		},
	}
	for _, c := range cases {
		p := NewParser([]byte(c.code))
		tokens, err := p.Parse()
		if err != nil {
			t.Fatalf("parse error %s", err)
		}
		for i, tp := range tokens {
			if tp.Type != c.types[i] {
				t.Errorf("expect %s but got %s", c.types[i], tp.Type)
			}
		}
	}
}
