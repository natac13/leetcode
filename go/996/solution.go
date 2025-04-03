package main

func subarraysWithKDistinct(nums []int, k int) int {
	count := make(map[int]int)
	result := 0
	leftFar := 0
	leftNear := 0

	for right := 0; right < len(nums); right++ {
		// add the current number to the count map
		count[nums[right]]++

		// shrink when count is greater than k
		for len(count) > k {
			count[nums[leftNear]]--
			if count[nums[leftNear]] == 0 {
				delete(count, nums[leftNear])
			}
			leftNear++

			leftFar = leftNear
		}

		// shrink when we have an abundance of a number in count
		for count[nums[leftNear]] > 1 {
			count[nums[leftNear]]--
			leftNear++
		}

		// if we have exactly k distinct numbers, we can count the number of subarrays
		// that can be formed with the current left and right pointers
		if len(count) == k {
			result += leftNear - leftFar + 1
		}

	}

	return result
}
