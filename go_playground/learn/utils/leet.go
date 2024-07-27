package utils

import (
	"math"
)
func FindClosetToZero(nums []int) int{
	closestNumber := math.MaxInt64 

	for _, num := range nums {
		if abs(num) < abs(closestNumber) || (abs(num) == abs(closestNumber) && num > 0 && closestNumber < 0) {
			closestNumber = num
		}
	}

	return closestNumber
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
