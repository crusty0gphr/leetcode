package reverse_number

import "testing"

var testNums = map[int]int{
	123: 321,
	67:  76,
	444: 444,
	2:   2,
	0:   0,
}

func TestReverseNumber(t *testing.T) {
	for n, r := range testNums {
		res := reverse(n)

		if r != res {
			t.Errorf("f.reverse(%v) | expects: %v | actual: %v", n, r, res)
		}
	}
}
