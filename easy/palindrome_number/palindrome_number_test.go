package palindrome_number

import "testing"

var testNum = map[int]bool{
	12:     false,
	121:    true,
	244212: false,
	1:      true,
	2332:   true,
}

func TestIsPalindromicNumber(t *testing.T) {
	for n, e := range testNum {
		r := isPalindrome(n)

		if e != r {
			t.Errorf("f.isPalindrome(%v) | expects: %v | actual: %v", n, e, r)
		}
	}
}
