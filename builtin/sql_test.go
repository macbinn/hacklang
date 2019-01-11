package builtin

import "testing"

func TestConvertName(t *testing.T) {
	cases := []struct{
		in string
		out string
	} {
		{"passwordHash", "password_hash"},
		{"name", "name"},
		{"userNameId", "user_name_id"},
	}
	for _, c := range cases {
		got := convertName(c.in)
		if got != c.out {
			t.Errorf("expect %s but got %s", c.out, got)
		}
	}
}

func TestConvertHackName(t *testing.T) {
	cases := []struct{
		in string
		out string
	} {
		{"password_hash", "passwordHash"},
		{"name", "name"},
		{"user_name_id", "userNameId"},
	}
	for _, c := range cases {
		got := convertHackName(c.in)
		if got != c.out {
			t.Errorf("expect %s but got %s", c.out, got)
		}
	}
}
