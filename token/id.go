package token

type idHandler struct {
}

func match(buf []byte, m string) bool {
	if len(buf) < len(m) {
		return false
	}
	return string(buf[:len(m)]) == m
}

var keywords = map[string]bool {
	"true": true,
	"false": true,
	"if": true,
	"and": true,
	"or": true,
	"not": true,
	"return": true,
}

func (idHandler) Match(buf []byte) (*Token, int, error) {
	pos := 0
	c := buf[pos]
	for (c >= 'a' && c < 'z') || (c >= 'A' && c <= 'Z') {
		pos++
		if pos >= len(buf) {
			break
		}
		c = buf[pos]
	}
	if pos == 0 {
		return nil, 0, ErrNotMatched
	}
	t := &Token{
		Type:  ID,
		Value: string(buf[0:pos]),
	}
	if keywords[t.Value] {
		return nil, 0, ErrNotMatched
	}
	return t, pos, nil
}

func init() {
	Register(string(ID), idHandler{})
}
