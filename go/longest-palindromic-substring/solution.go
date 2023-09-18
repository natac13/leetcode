package longestpalindromicsubstring

// longestPalindrome returns the longest palindromic substring of the input string s.
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	var low, maxLen int

	// Iterate over each character in the string.
	for i := 0; i < len(s); i++ {
		// Find the longest palindrome with odd length that has s[i] as its center.
		low1, len1 := expandFromCenter(s, i, i)
		// Find the longest palindrome with even length that has s[i] and s[i+1] as its center.
		low2, len2 := expandFromCenter(s, i, i+1)

		// If the palindrome with even length is longer than the palindrome with odd length
		// and it's longer than the current maximum length, update the maximum length and the starting index.
		if len2 > len1 && len2 > maxLen {
			maxLen = len2
			low = low2
			// Otherwise, if the palindrome with odd length is longer than the current maximum length,
			// update the maximum length and the starting index.
		} else if len1 > maxLen {
			maxLen = len1
			low = low1
		}

	}

	// Return the longest palindromic substring.
	return s[low : low+maxLen+1]

}

// expandFromCenter returns the length and starting index of the longest palindrome
// that has s[i] and s[j] as its center.
func expandFromCenter(s string, left, right int) (int, int) {

	// making sure left is not less than 0 and right is not greater than the length of the string
	// and s[left] == s[right]
	// while the above conditions are true, decrement left and increment right
	// this will expand the potential palindrome out from the center
	// when the conditions are no longer true, return the starting index and the length of the palindrome
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}

	// return the starting index and the length of the palindrome
	// the length of the palindrome is right - left - 1 because
	// left and right are no longer the center of the palindrome
	left += 1
	right -= 1

	return left, right - left
}

func longestPalindromeV2(s string) string {

	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // length of odd palindrome
		len2 := expandAroundCenter(s, i, i+1) // length of even palindrome
		len := max(len1, len2)
		if len > end-start+1 {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l := left
	r := right

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l -= 1
		r += 1
	}

	return r - l - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
