package format

import "github.com/macbinn/hacklang/parser"

func Format(file string) (string, error) {
	node, err := parser.ParseFile(file)
	if err != nil {
		return "", err
	}
	return node.Code(), nil
}
