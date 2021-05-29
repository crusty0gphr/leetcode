package main

import "fmt"

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToIntRec(s string, res int) int {
	chars := []byte(s)

	if len(chars) >= 2 {
		num := roman[string(chars[0])]
		next := roman[string(chars[1])]

		if next > num {
			res = res - num
		} else {
			res = res + num
		}
		s = string(chars[1:]) // pop chars[0] and convert ot string
		return romanToIntRec(s, res)
	}

	if len(chars) == 1 {
		return res + roman[string(chars[0])]
	}

	return res
}

func main() {
	iv := romanToIntRec("IV", 0)
	iii := romanToIntRec("III", 0)
	mdxiii := romanToIntRec("MDXIII", 0)
	mdlix := romanToIntRec("MDLXV", 0)
	ix := romanToIntRec("IX", 0)

	fmt.Println(iv, iii, mdxiii, mdlix, ix)
}
