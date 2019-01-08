package token

type boolHandler struct {
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func matchSpace(buf []byte, idx int) bool {
	if idx >= len(buf) {
		return true
	}
	if buf[idx] >= 'a' && buf[idx] <= 'z' {
		return false
	}
	if buf[idx] >= 'A' && buf[idx] <= 'Z' {
		return false
	}
	return true
}

func (boolHandler) Match(buf []byte) (*Token, int, error) {
	if match(buf, TRUE) && matchSpace(buf, 4) {
		t := &Token{
			Type:  BOOL,
			Value: TRUE,
		}
		return t, 4, nil
	}
	if match(buf, FALSE) && matchSpace(buf, 5) {
		t := &Token{
			Type:  BOOL,
			Value: FALSE,
		}
		return t, 5, nil
	}
	return nil, 0, ErrNotMatched
}

func init() {
	Register(string(BOOL), boolHandler{})
}
