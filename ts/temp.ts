import * as assert from 'node:assert'

const fields = [
	{
		id: 'A',
		name: 'A',
	},
	{
		id: 'B',
		name: 'B',
	},
	{
		id: 'C',
		name: 'C',
	},
	{
		id: 'D',
		name: 'D',
	},
	{
		id: 'E',
		name: 'E',
	},
	{
		id: 'F',
		name: 'F',
	},
	{
		id: 'G',
		name: 'G',
	},
]

function moveFieldAfter(fields: any[], fieldId: string, afterFieldId: string) {
	const fieldIdx = fields.findIndex((f) => f.id === fieldId)
	if (fieldIdx === -1) return false // Handle missing field

	// special case: if afterFieldId is 0, move to start of array
	if (afterFieldId === '0') {
		const field = fields[fieldIdx]
		fields.splice(fieldIdx, 1) // Remove field
		fields.unshift(field) // Insert at start of array
		return true
	}

	const afterIdx = fields.findIndex((f) => f.id === afterFieldId)
	if (afterIdx === -1) return false // Handle missing field

	if (fieldId === afterFieldId) return true // No-op
	if (fieldIdx === afterIdx + 1) return true // No-op

	const field = fields[fieldIdx]
	fields.splice(fieldIdx, 1) // Remove field

	// If field was before insertion point, the insertion point shifted down by 1
	const insertAt = fieldIdx < afterIdx ? afterIdx : afterIdx + 1
	fields.splice(insertAt, 0, field) // Insert after target field

	return true
}

function main() {
	const cases = [
		{
			fieldId: 'C',
			afterFieldId: '0',
			result: true,
			expected: [
				{ id: 'C', name: 'C' },
				{ id: 'A', name: 'A' },
				{ id: 'B', name: 'B' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
				{ id: 'G', name: 'G' },
			],
		},
		{
			fieldId: 'X',
			afterFieldId: '0',
			result: false,
			expected: [
				{ id: 'A', name: 'A' },
				{ id: 'B', name: 'B' },
				{ id: 'C', name: 'C' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
				{ id: 'G', name: 'G' },
			],
		},
		{
			fieldId: 'A',
			afterFieldId: 'B',
			result: true,
			expected: [
				{ id: 'B', name: 'B' },
				{ id: 'A', name: 'A' },
				{ id: 'C', name: 'C' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
				{ id: 'G', name: 'G' },
			],
		},
		{
			fieldId: 'A',
			afterFieldId: 'C',
			result: true,
			expected: [
				{ id: 'B', name: 'B' },
				{ id: 'C', name: 'C' },
				{ id: 'A', name: 'A' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
				{ id: 'G', name: 'G' },
			],
		},
		{
			fieldId: 'G',
			afterFieldId: 'F',
			result: true,
			expected: [
				{ id: 'A', name: 'A' },
				{ id: 'B', name: 'B' },
				{ id: 'C', name: 'C' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
				{ id: 'G', name: 'G' },
			],
		},
		{
			fieldId: 'G',
			afterFieldId: 'A',
			result: true,
			expected: [
				{ id: 'A', name: 'A' },
				{ id: 'G', name: 'G' },
				{ id: 'B', name: 'B' },
				{ id: 'C', name: 'C' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
			],
		},
		{
			fieldId: 'E',
			afterFieldId: 'C',
			result: true,
			expected: [
				{ id: 'A', name: 'A' },
				{ id: 'B', name: 'B' },
				{ id: 'C', name: 'C' },
				{ id: 'E', name: 'E' },
				{ id: 'D', name: 'D' },
				{ id: 'F', name: 'F' },
				{ id: 'G', name: 'G' },
			],
		},

		{
			fieldId: 'G',
			afterFieldId: 'C',
			result: true,
			expected: [
				{ id: 'A', name: 'A' },
				{ id: 'B', name: 'B' },
				{ id: 'C', name: 'C' },
				{ id: 'G', name: 'G' },
				{ id: 'D', name: 'D' },
				{ id: 'E', name: 'E' },
				{ id: 'F', name: 'F' },
			],
		},
	]

	for (const testCase of cases) {
		const { fieldId, afterFieldId, expected } = testCase
		const copy = fields.slice()
		const result = moveFieldAfter(copy, fieldId, afterFieldId)

		assert.equal(
			result,
			testCase.result,
			`Expected ${fieldId} to be moved after ${afterFieldId}`,
		)

		assert.deepEqual(
			copy,
			expected,
			`Expected ${JSON.stringify(copy)} to be ${JSON.stringify(expected)}`,
		)
	}
}

main()
