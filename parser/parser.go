package parser

import (
	"fmt"
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type Parser struct {
	code   []byte
	tokens []*token.Token
	pos    int
}

type Handler interface {
	Parse(tokens []*token.Token) (ast.Node, int, error)
}

var handlers = map[string]Handler{}
var (
	ErrHandlerNotFound = fmt.Errorf("handler not found")
	ErrSyntaxError     = fmt.Errorf("synatx error")
)

func Register(name string, handler Handler) {
	handlers[name] = handler
}

func Parse(name string, tokens []*token.Token) (ast.Node, int, error) {
	handler, ok := handlers[name]
	if !ok {
		return nil, 0, ErrHandlerNotFound
	}
	return handler.Parse(tokens)
}

func ParseGreedy(tokens []*token.Token, names ...string) (node ast.Node, pos int, err error) {
	if len(tokens) == 0 {
		return nil, 0, ErrSyntaxError
	}
	for _, name := range names {
		n, i, err := Parse(name, tokens)
		if err == nil {
			if i > pos {
				pos = i
				node = n
			}
		}
	}
	if node != nil {
		return node, pos, nil
	}
	return nil, 0, ErrSyntaxError
}

func NewParser(code []byte) *Parser {
	return &Parser{
		code: code,
	}
}

func (p *Parser) Parse() (ast.Node, error) {
	tokenParser := token.NewParser(p.code)
	tokens, err := tokenParser.Parse()
	if err != nil {
		return nil, err
	}
	p.tokens = tokens
	var names []string
	for name := range handlers {
		names = append(names, name)
	}
	node, _, err := ParseGreedy(p.tokens, names...)
	if err != nil {
		return nil, ErrSyntaxError
	}
	return node, nil
}
