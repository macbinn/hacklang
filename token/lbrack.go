package token

type staticStringHandler struct {
	Matched string
	Type    Type
}

func (h staticStringHandler) Match(buf []byte) (*Token, int, error) {
	if match(buf, h.Matched) {
		t := &Token{
			Type:  h.Type,
			Value: string(h.Type),
		}
		return t, 1, nil
	}
	return nil, 0, ErrNotMatched
}

func init() {
	Register(string(LBRACK), staticStringHandler{"[", LBRACK})
	Register(string(RBRACK), staticStringHandler{"]", RBRACK})
	Register(string(COLON), staticStringHandler{":", COLON})
	Register(string(COMMA), staticStringHandler{",", COMMA})
	Register(string(LBRACE), staticStringHandler{"{", LBRACE})
	Register(string(RBRACE), staticStringHandler{"}", RBRACE})
	Register(string(EQUAL), staticStringHandler{"=", EQUAL})
	Register(string(LPAREN), staticStringHandler{"(", LPAREN})
	Register(string(RPAREN), staticStringHandler{")", RPAREN})
	Register(string(RARROW), staticStringHandler{">", RARROW})
	Register(string(DOT), staticStringHandler{".", DOT})
}
