package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func FirstNonRepeating(str string) string {
	if len(str) <= 1 {
		return str
	}
	repeatedStr := map[string]string{}
	for i := 0; i < len(str); i++ {
		char := str[i]
		answer := string(char)
		for j := i + 1; j < len(str); j++{
			if repeatedStr[string(char)] == string(char){
				answer = ""
			}
			if strings.EqualFold(string(char), string(str[j])) {
				answer = ""
				repeatedStr[string(char)] = string(char)
			}
		}
		if repeatedStr[answer] != answer{
			return answer
		}
	}
	return ""
}

// ".*?" ==> To(Equal(false)) ✔️
	// "a" ==> To(Equal(true)) ✔️
	// "Mazinkaiser" ==> To(Equal(true)) ✔️
	// "hello world_" ==> To(Equal(false)) ✔️
	// "     " ==> To(Equal(false)) ✔️
	// "PassW0rd" ==> To(Equal(true)) ✔️
	// "" ==> To(Equal(false)) ✔️
	// "\n\t\n" ==> To(Equal(false)) ✔️
	// "ciao\n$$_" ==> To(Equal(false)) ✔️
	// "__ * __" ==> To(Equal(false)) ✔️
	// "&)))(((" ==> To(Equal(false)) ✔️
	// "43534h56jmTHHF3k" ==> To(Equal(true)) ✔️

func Alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, val := range str {
		switch {
		case val < 48 || val > 57 && val < 65 || val > 90 && val < 97:
			return false
		default:
			continue
		}
	}

	return true
}

func WordCount(s string) map[string]int {
	countMap := map[string]int{}
	stringArr := strings.Split(s, " ")
	for _, word := range stringArr {
	w, ok := countMap[word]
	 if ok {
	 	countMap[word] = w + 1
		continue
	 }
	 countMap[word] = 1
	}
	return countMap
}