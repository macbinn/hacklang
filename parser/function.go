package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type functionHandler struct {
}

func parseParamList(tokens []*token.Token) ([]string, int, error) {
	if tokens[0].Type == token.ID {
		return []string{tokens[0].Value}, 1, nil
	}
	if tokens[0].Type != token.LPAREN {
		return nil, 0, ErrSyntaxError
	}
	pos := 1
	var params []string
	for pos < len(tokens) {
		if tokens[pos].Type == token.ID {
			params = append(params, tokens[pos].Value)
		} else {
			return nil, 0, ErrSyntaxError
		}
		pos ++
		if pos >= len(tokens) {
			return nil, 0, ErrSyntaxError
		}
		if tokens[pos].Type == token.COMMA {
			pos ++
			continue
		} else if tokens[pos].Type == token.RPAREN {
			return params, pos + 1, nil
		}
	}
	return nil, 0, ErrSyntaxError
}

func (functionHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	params, pos, err := parseParamList(tokens)
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	if pos + 3 > len(tokens) ||
		tokens[pos].Type != token.EQUAL ||
		tokens[pos + 1].Type != token.RARROW ||
		tokens[pos + 2].Type != token.LBRACE {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.FunctionNode{
		Arguments: params,
	}
	pos += 3
	if tokens[pos].Type == token.RBRACE {
		pos++
		return node, pos, nil
	}
	body, i, err := Parse("exprList", tokens[pos:])
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	node.Body = body
	pos += i
	if pos < len(tokens) && tokens[pos].Type == token.RBRACE {
		pos++
		return node, pos, nil
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("function", functionHandler{})
}
