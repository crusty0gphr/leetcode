package roman_to_integer_loop

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	res := 0

	cArr := []byte(s)
	cLen := len(cArr)

	for i := cLen - 1; i >= 0; i-- {
		r := string(cArr[i])
		rEqw := roman[r]

		if i+1 < cLen {
			rNext := string(cArr[i+1])
			rEqeNext := roman[rNext]

			if rEqw >= rEqeNext {
				res = res + rEqw
			} else {
				res = res - rEqw
			}
		} else {
			res = res + rEqw
		}
	}

	return res
}
