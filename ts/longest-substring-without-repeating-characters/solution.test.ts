import { describe, expect, it } from 'vitest'

import { lengthOfLongestSubstring } from './solution.ts'

describe('Longest Substring Without Repeating Characters', () => {
	it('should return 3 when input is "abcabcbb"', () => {
		const input = 'abcabcbb'
		const output = lengthOfLongestSubstring(input)
		expect(output).toEqual(3)
	})

	it('should return 1 when input is "bbbbb"', () => {
		const input = 'bbbbb'
		const output = lengthOfLongestSubstring(input)
		expect(output).toEqual(1)
	})

	it('should return 3 when input is "pwwkew"', () => {
		const input = 'pwwkew'
		const output = lengthOfLongestSubstring(input)
		expect(output).toEqual(3)
	})

	it('should return 0 when input is ""', () => {
		const input = ''
		const output = lengthOfLongestSubstring(input)
		expect(output).toEqual(0)
	})
})
