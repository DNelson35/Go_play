package main

import (
	"fmt"
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

var counts = WordCount("the one who knows one is one who knows none")

func main(){
	fmt.Println(counts)
}