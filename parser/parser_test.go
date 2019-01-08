package parser

import "testing"

func TestParse(t *testing.T) {
	cases := []string{
		"a = 123",
		"a = true",
		"a = false",
		"a = `abc`",
		"a = [1, 2, 3]",
		"a = f()",
		"f(i)",
		"i => {}",
		"i => {f()}",
		`a = i => {
   print(i)
}`,
	}
	for _, c := range cases {
		p := NewParser([]byte(c))
		_, err := p.Parse()
		if err != nil {
			t.Errorf("%s: parse error %s", c, err)
		}
	}
}
