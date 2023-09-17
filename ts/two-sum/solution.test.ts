import { describe, expect, it } from 'vitest'

import { twoSum } from './solution.ts'

describe('twoSum', () => {
	it('should return the indices of the two numbers that add up to the target', () => {
		const nums = [2, 7, 11, 15]
		const target = 9
		const expected = [0, 1]

		expect(twoSum(nums, target)).toEqual(expected)
	})

	it('should return an empty array if no two numbers add up to the target', () => {
		const nums = [2, 7, 11, 15]
		const target = 10

		expect(twoSum(nums, target)).toEqual([])
	})

	it('should handle negative numbers', () => {
		const nums = [-2, 7, 11, -15]
		const target = -17
		const expected = [0, 3]

		expect(twoSum(nums, target)).toEqual(expected)
	})

	it('should handle duplicate numbers', () => {
		const nums = [3, 2, 4]
		const target = 6
		const expected = [1, 2]

		expect(twoSum(nums, target)).toEqual(expected)
	})
})
