package twosum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}

	return []int{}
}
