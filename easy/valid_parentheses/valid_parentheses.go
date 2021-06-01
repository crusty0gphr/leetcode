package valid_parentheses

var pairs = map[string]string{"}": "{", ")": "(", "]": "["}

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(v string) {
	*s = append(*s, v)
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		i := len(*s) - 1
		ppd := (*s)[i]
		*s = (*s)[:i]
		return ppd, true
	}
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	var stack Stack
	for _, n := range s {
		c := string(n)
		if _, ok := pairs[c]; ok {
			ppd, _ := stack.pop()
			if pairs[c] != ppd {
				return false
			}
		} else {
			stack.push(c)
		}
	}
	return stack.isEmpty()
}
