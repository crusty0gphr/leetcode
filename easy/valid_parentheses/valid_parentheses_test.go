package valid_parentheses

import (
	"fmt"
	"testing"
)

var except = []bool{
	true,
	true,
	false,
	true,
	false,
	true,
	false,
	false,
}

var parentheses = []string{
	"(()(({})[]))",
	"()()()()",
	"{}[][[[())",
	"{(([()](){}))}",
	"",
	"()",
	"(]",
	"(]]",
}

func TestIsValid(t *testing.T) {
	for i := 0; i < len(except); i++ {
		e := except[i]
		v := parentheses[i]

		res := isValid(v)

		fmt.Printf("[OUTPUT - %v] v:%v, e:%v, a:%v\n", i, v, e, res)
		if res != e {
			t.Errorf("f.isValid(%v) | expects: %v | actual: %v", v, e, res)
		}
	}
}
