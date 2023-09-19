package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// reverse the number
	reversed := 0
	// store the original number
	original := x
	// while x is greater than 0
	for x > 0 {
		// multiply reversed by 10 to make room for the next digit
		carry := reversed * 10
		// add the last digit of x to reversed
		// x % 10 gets the last digit of x
		addition := x % 10
		// add the last digit of x to reversed
		// and store it in reversed
		reversed = carry + addition
		// remove the last digit of x
		x /= 10
	}

	return original == reversed
}
