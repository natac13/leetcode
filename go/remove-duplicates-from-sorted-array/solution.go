package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	start := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			temp := nums[i]
			nums[start] = temp
			start += 1
		}
	}

	return start
}
