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
		return t, len(h.Matched), nil
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
	Register(string(EQUALS), staticStringHandler{"==", EQUALS})
	Register(string(EQUAL), staticStringHandler{"=", EQUAL})
	Register(string(LPAREN), staticStringHandler{"(", LPAREN})
	Register(string(RPAREN), staticStringHandler{")", RPAREN})
	Register(string(RARROW), staticStringHandler{">", RARROW})
	Register(string(DOT), staticStringHandler{".", DOT})
	Register(string(IF), staticStringHandler{"if", IF})
	Register(string(AND), staticStringHandler{"and", AND})
	Register(string(OR), staticStringHandler{"or", OR})
	Register(string(NOT), staticStringHandler{"not", NOT})
	Register(string(RETURN), staticStringHandler{"return", RETURN})
	Register(string(PLUS), staticStringHandler{"+", PLUS})
	Register(string(MINS), staticStringHandler{"-", MINS})
	Register(string(MUL), staticStringHandler{"*", MUL})
	Register(string(DEV), staticStringHandler{"/", DEV})
}
