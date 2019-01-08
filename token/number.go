package token

type numberHandler struct {
}

func (numberHandler) Match(buf []byte) (*Token, int, error) {
	pos := 0
	c := buf[pos]
	for c >= '0' && c <= '9' {
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
		Type:  NUMBER,
		Value: string(buf[0:pos]),
	}
	return t, pos, nil
}

func init() {
	Register(string(NUMBER), numberHandler{})
}
