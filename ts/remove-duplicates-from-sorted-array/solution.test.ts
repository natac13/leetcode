import { describe, expect, it } from 'vitest'

import { removeDuplicates } from './solution.ts'

describe('removeDuplicates', () => {
	it.only('should return the length of the array with duplicates removed', () => {
		// expect(removeDuplicates([1, 1, 2])).toBe(2)
		expect(removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])).toBe(5)
		// expect(removeDuplicates([1, 2, 3])).toBe(3)
		// expect(removeDuplicates([1, 1, 1, 1, 1])).toBe(1)
	})

	it('should modify the input array to have the first n unique elements', () => {
		const nums1 = [1, 1, 2]
		expect(removeDuplicates(nums1)).toBe(2)
		expect(nums1.slice(0, 2)).toEqual([1, 2])

		const nums2 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
		expect(removeDuplicates(nums2)).toBe(5)
		expect(nums2.slice(0, 5)).toEqual([0, 1, 2, 3, 4])

		const nums3 = [1, 2, 3]
		expect(removeDuplicates(nums3)).toBe(3)
		expect(nums3).toEqual([1, 2, 3])
	})
})
