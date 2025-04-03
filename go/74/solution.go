package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	if m == 0 || n == 0 {
		return false
	}

	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return false
}
