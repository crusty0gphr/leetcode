package roman_to_integer_loop

import "testing"

var testRomans = map[string]int{
	"IV":     4,
	"MDLII":  1552,
	"III":    3,
	"IX":     9,
	"MDXIII": 1513,
}

func TestRomanToInteger(t *testing.T) {
	for s, e := range testRomans {
		res := romanToInt(s)

		if res != e {
			t.Errorf("f.romanToInt(%v) | expects: %v | actual: %v", s, e, res)
		}
	}
}
