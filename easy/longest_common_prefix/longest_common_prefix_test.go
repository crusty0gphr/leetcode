package longest_common_prefix

import (
	"fmt"
	"testing"
)

var testArgs = map[int][]string{
	0: {"fl", "flower", "flow", "flight"},
	1: {"", "dog", "racecar", "car"},
	2: {"dod", "dodge", "dodo", "dodenic"},
	3: {""},
	4: {"pog", "pog"},
	5: {"a", "ab", "a"},
	6: {"aa", "aaa", "aa", "aaa"},
	7: {"", "fun", "pikachu", "meteor", "ska"},
}

func TestLongestCommonPrefix(t *testing.T) {
	for i := 0; i < len(testArgs); i++ {
		e := testArgs[i][0]
		r := longestCommonPrefix(testArgs[i][1:])

		if e != r {
			t.Errorf("f.longestCommonPrefix(%v) | expects: %v | actual: %v", testArgs[i][1:], e, r)
		}

		fmt.Printf("--- Output [%d]: %v \n", i, r)
	}
}
