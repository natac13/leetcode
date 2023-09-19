export function removeDuplicates(nums: number[]): number {
	// start at 1 because the first element is always unique
	let start = 1

	// iterate through the array
	for (let i = 1; i < nums.length; i++) {
		const curr = nums[i]
		const prev = nums[i - 1]
		// if the current element is different from the previous element
		// and the current element is not undefined
		if (curr !== prev && typeof curr !== 'undefined') {
			// set the element at the start index to the current element
			nums[start] = curr
			// increment the start index
			start++
		}
	}

	return start
}
