package token

import (
	"fmt"
)

type Type string

type Token struct {
	Type  Type
	Value string
}

const (
	ID     Type = "id"     // a
	NUMBER Type = "number" // 12345
	STRING Type = "string" // `abc`
	BOOL   Type = "bool"   // true
	LBRACK Type = "["
	RBRACK Type = "]"
	LBRACE Type = "{"
	RBRACE Type = "}"
	COMMA  Type = ","
	COLON  Type = ":"
	EQUAL  Type = "="
	LPAREN Type = "("
	RPAREN Type = ")"
	RARROW Type = ">"
	DOT    Type = "."
)

type Parser struct {
	buf []byte
	pos int
}

var (
	ErrNotMatched = fmt.Errorf("not matched")
	ErrNotToken   = fmt.Errorf("not token")
)

func NewParser(buf []byte) *Parser {
	return &Parser{
		buf: buf,
		pos: 0,
	}
}

type Handler interface {
	Match(buf []byte) (*Token, int, error)
}

var handlers = map[string]Handler{}

func Register(name string, handler Handler) {
	handlers[name] = handler
}

var ErrEOF = fmt.Errorf("EOF")

func (p *Parser) NextToken() (*Token, error) {
	if p.pos >= len(p.buf) {
		return nil, ErrEOF
	}
	c := p.buf[p.pos]
	for c == ' ' || c == '\n' {
		p.pos++
		if p.pos >= len(p.buf) {
			return nil, ErrEOF
		}
		c = p.buf[p.pos]
	}
	for _, handler := range handlers {
		token, pos, err := handler.Match(p.buf[p.pos:])
		if err == nil {
			p.pos += pos
			return token, nil
		}
	}
	return nil, ErrNotToken
}

func (p *Parser) Parse() (tokens []*Token, err error) {
	for {
		t, err := p.NextToken()
		if err == ErrEOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	return tokens, nil
}
