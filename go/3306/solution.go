package main

func countOfSubstrings(word string, k int) int64 {

	return atLeastK(word, k) - atLeastK(word, k+1)

}

func atLeastK(word string, k int) int64 {
	vowels := map[rune]int{}
	non_vowel_count := 0
	result := 0
	left := 0

	for right, c := range word {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels[c] += 1
		} else {
			non_vowel_count++
		}

		vowels_count := len(vowels)

		for vowels_count == 5 && non_vowel_count >= k {
			result += (len(word) - right)

			leftChar := rune(word[left])
			if leftChar == 'a' || leftChar == 'e' || leftChar == 'i' || leftChar == 'o' || leftChar == 'u' {
				vowels[leftChar]--
				if vowels[leftChar] == 0 {
					delete(vowels, leftChar)
					vowels_count--
				}

			} else {
				non_vowel_count--
			}
			left++
		}
	}

	return int64(result)

}
