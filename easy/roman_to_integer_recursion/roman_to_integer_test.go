package roman_to_integer_recursion

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
		res := romanToInt(s, 0)

		if res != e {
			t.Errorf("f.romanToInt(%v) | expects: %v | actual: %v", s, e, res)
		}
	}
}
