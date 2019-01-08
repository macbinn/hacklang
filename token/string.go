package token

type stringHandler struct {
}

func (stringHandler) Match(buf []byte) (*Token, int, error) {
	if buf[0] != '`' {
		return nil, 0, ErrNotMatched
	}
	pos := 1
	end := false
	for pos < len(buf) {
		c := buf[pos]
		if c == '`' {
			end = true
			break
		}
		pos++
	}
	if !end {
		return nil, 0, ErrNotMatched
	}
	t := &Token{
		Type:  STRING,
		Value: string(buf[1:pos]),
	}
	return t, pos + 1, nil
}

func init() {
	Register(string(STRING), stringHandler{})
}
