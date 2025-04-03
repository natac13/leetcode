package main

import "math"

func checkPowersOfThree(n int) bool {
	i := 0
	for ; int(math.Pow(3, float64(i))) < n; i++ {
	}

	i--

	for i >= 0 {
		power := int(math.Pow(3, float64(i)))
		if power <= n {
			n -= power
		}
		if power <= n {
			return false
		}
		i--
	}

	return n == 0

}
