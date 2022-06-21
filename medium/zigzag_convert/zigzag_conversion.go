package zigzagconvert

import (
	"strings"
)

func convert(s string, n int) string {
	if n == 1 {
		return s
	}

	r := make([]string, n)

	var p int
	var d bool

	for i := range s {
		r[p] += string(s[i])
		if p == 0 || p == n-1 {
			d = !d
		}

		if d {
			p++
		} else {
			p--
		}
	}

	return strings.Join(r, "")
}

func convertMatrix(s string, n int) string {
	if n == 1 {
		return s
	}

	var (
		b   = make([][]byte, n)
		r   int
		inc = 1
	)

	for i := 0; i < len(s); i++ {
		b[r] = append(b[r], s[i])
		r += inc
		if r == n-1 || r == 0 {
			inc *= -1
		}
	}

	var sb strings.Builder
	for _, v := range b {
		sb.Write(v)
	}

	return sb.String()
}
