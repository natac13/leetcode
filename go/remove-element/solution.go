package removeelement

func removeElement(nums []int, val int) int {
	res := 0

	lIdx := 0
	rIdx := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if lIdx > rIdx {
			break
		}

		if lIdx == rIdx {
			if nums[lIdx] != val {
				res += 1
			}
			break
		}

		lItem := nums[lIdx]
		rItem := nums[rIdx]

		// when both left and right equal val then move right pointer
		// this will be handled by the next iteration which finds a right pointer that is not equal to val
		if rItem == val && lItem == val {
			nums[rIdx] = -1
			rIdx -= 1
			continue
		}

		// when right is equal to val and left is not equal to val then move both pointers
		// and increment result
		if rItem == val && lItem != val {
			nums[rIdx] = -1
			rIdx -= 1
			lIdx += 1
			res += 1
			continue
		}

		// when both left and right are not equal to val then move only left pointer
		// and increment result by 1
		if rItem != val && lItem != val {
			lIdx += 1
			res += 1
			continue
		}

		// when left equals val and right does not equal val then swap left and right
		// and increment result by 1
		if lItem == val && rItem != val {
			temp := nums[rIdx]
			nums[lIdx] = temp
			nums[rIdx] = -1
			lIdx += 1
			rIdx -= 1
			res += 1
			continue
		}

	}

	return res
}

func removeElementV2(nums []int, val int) int {
	j := 0

	for _, v := range nums {
		if v != val {
			nums[j] = v
			j += 1
		}
	}
	return j
}
