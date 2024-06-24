package utils

import (
	"strings"
)

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