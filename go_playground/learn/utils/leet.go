package utils

import (
	"math"
	"strconv"
	"fmt"
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

// 56789 -> 68957
func MaxRot(n int64) int64 {
    var highValue int64
	numStr := strconv.FormatInt(n, 10)
	for i := 0; i < len(numStr)-1; i++{
		if i != len(numStr) - 1{
			firstChunk := numStr[:i]
			secondChunk := numStr[i + 1:]
			numStr = firstChunk + secondChunk + string(numStr[i:i+1])
			fmt.Println(numStr)
		}
		newNum, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			fmt.Println("not a number")
		}
		if highValue < newNum{
			highValue = newNum
		}
	}
	return highValue
}
