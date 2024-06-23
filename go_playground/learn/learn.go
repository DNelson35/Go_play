package main

import (
	"fmt"
	// "strings"
)


// func WordCount(s string) map[string]int {
// 	countMap := map[string]int{}
// 	stringArr := strings.Split(s, " ")
// 	for _, word := range stringArr {
// 	w, ok := countMap[word]
// 	 if ok {
// 	 	countMap[word] = w + 1
// 		continue
// 	 }
// 	 countMap[word] = 1
// 	}
// 	return countMap
// }

// var counts = WordCount("the one who knows one is one who knows none")

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(x int) int {
	pos1, pos2 := 0,0

	return func(x int) int {
		if x == 1 {
			pos2 = 1
			return pos1 + pos2
		}else if x < 1 {
			return 0
		}
		num := pos1 + pos2
		pos1, pos2 = pos2, num
		return num
	}
}

func main(){
	// fmt.Println(counts)
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}