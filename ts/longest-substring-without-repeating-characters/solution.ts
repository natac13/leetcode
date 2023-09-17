export function lengthOfLongestSubstring(s: string): number {
	let max = 0
	let start = 0

	const map = new Map<string, number>()

	for (let i = 0; i < s.length; i++) {
		const char = s[i]

		if (!char) {
			continue
		}

		if (map.has(char)) {
			// If the character is already in the map, we need to update the start
			// position to the next character after the last time we saw this
			// character.
			start = Math.max(map.get(char)! + 1, start)
		}

		// Update the map with the current character and its index.
		map.set(char, i)
		// Update the max length if the current length is greater than the max.
		max = Math.max(i - start + 1, max)
	}

	return max
}
