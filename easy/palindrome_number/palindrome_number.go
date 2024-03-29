package palindrome_number

func isPalindrome(x int) bool {
	isPalindrome := false
	revX := reverseInt(x)

	if revX == x {
		isPalindrome = true
	}

	return isPalindrome
}

func reverseInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}
