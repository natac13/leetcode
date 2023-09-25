package applyoperationstoanarray

func applyOperations(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j += 1
		}
	}

	return nums
}

func applyOperationsV0(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}

	j := 0

	for i, v := range nums {
		if v != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j += 1
		}
	}

	return nums
}
