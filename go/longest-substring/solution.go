package longestsubstring

// lengthOfLongestSubstring returns the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	// max keeps track of the maximum length of the substring without repeating characters.
	// start keeps track of the starting index of the substring without repeating characters.
	max, start := 0, 0

	// seen is a map that keeps track of the last index of each character seen in the string.
	seen := make(map[rune]int)

	// Iterate over each character in the string.
	for i, c := range s {
		// If the character has been seen before and its last index is greater than or equal to the starting index of the current substring,
		// update the starting index of the current substring to the index after the last occurrence of the character.
		if prev, ok := seen[c]; ok && prev >= start {
			start = prev + 1
		}

		// Update the last index of the character to the current index.
		seen[c] = i

		// Update the maximum length of the substring without repeating characters if the current substring is longer.
		currLen := i - start + 1
		if currLen > max {
			max = currLen
		}
	}

	// Return the maximum length of the substring without repeating characters.
	return max
}
