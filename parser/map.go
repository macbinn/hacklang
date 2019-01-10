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
	if tokens[0].Type == token.LBRACE && tokens[1].Type == token.RBRACE {
		node := &ast.MapNode{}
		return node, 2, nil
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("map", mapHandler{})
}

