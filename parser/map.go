package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type mapHandler struct {

}

func (mapHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 2 {
		return nil, 0, ErrSyntaxError
	}
	if tokens[0].Type != token.LBRACE {
		return nil, 0, ErrSyntaxError
	}
	pos := 1
	node := &ast.MapNode{
		Init: map[string]ast.Node{},
	}
	for pos < len(tokens) {
		if tokens[pos].Type == token.RBRACE {
			return node, pos + 1, nil
		}
		if tokens[pos].Type != token.ID || tokens[pos + 1].Type != token.COLON {
			break
		}
		name := tokens[pos].Value
		pos += 2
		expr, i, err := ParseGreedy(tokens[pos:], "expr")
		if err != nil {
			break
		}
		pos += i
		node.Init[name] = expr
		if tokens[pos].Type == token.COMMA {
			pos ++
			continue
		} else if tokens[pos].Type == token.RBRACE {
			return node, pos + 1, nil
		}
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("map", mapHandler{})
}

