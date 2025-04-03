package main

func mergeArrays(nums1, nums2 [][]int) [][]int {
	var i, j int
	var result [][]int

	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] < nums2[j][0] {
			result = append(result, nums1[i])
			i++
		} else if nums1[i][0] > nums2[j][0] {
			result = append(result, nums2[j])
			j++
		} else {
			result = append(result, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		}
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}
	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}

	return result
}
