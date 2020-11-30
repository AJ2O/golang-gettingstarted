package morestrings

import "testing"

// No need to import this test file anywhere, just run "go test" in
// this directory

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	{
		for _, c := range cases {
			got := ReverseRunes(c.in)
			if got != c.want {
				t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}
