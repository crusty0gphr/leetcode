package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
    return reverseNumber(x, 0)
}

func reverseNumber(x int, rev int) int {
	pop := x % 10
	rev = rev*10 + pop

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	if pop != x {
		tX := x / 10
		return reverseNumber(tX, rev)
	}

	return rev
}

func main() {
	res := reverse(235)
	fmt.Println(res)
}
