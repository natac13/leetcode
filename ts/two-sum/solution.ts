export function twoSum(nums: number[], target: number): number[] {
	// where the key is the value and the value is the index in the array
	const seen: { [key: number]: number } = {}

	for (const [idx, currentNumber] of nums.entries()) {
		const diff = target - currentNumber

		const seenDiff = seen[diff]
		if (seenDiff !== undefined) {
			return [seenDiff, idx]
		}

		seen[currentNumber] = idx
	}

	return []
}
