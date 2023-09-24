package movezeros

func moveZeros(nums []int) {
	lastZeroIdx := -1

	for i, v := range nums {
		// if we have seen a zero prior to i and the current v is not a zero
		// SWAP
		if v != 0 && lastZeroIdx >= 0 {

			nums[lastZeroIdx], nums[i] = v, 0
			lastZeroIdx += 1
			// else if we have found our first zero
			// adjust the lastZeroIdx
		} else if lastZeroIdx < 0 && v == 0 {
			lastZeroIdx = i
		}
	}
}

func moveZerosV2(nums []int) {
	lastZeroIdx := 0

	for i, v := range nums {
		if v != 0 {
			nums[i], nums[lastZeroIdx] = nums[lastZeroIdx], nums[i]
			lastZeroIdx += 1
		}
	}
}
