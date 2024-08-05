package main

import (
	"fmt"
	"learn/utils"
)


func main(){
	

	// fmt.Println(utils.LgPalindrom())
	// fmt.Println(utils.FindDiffSquares(100))

	// fmt.Println(utils.GetPrime(10001))

	count := utils.WordCount("the one who knows one is one who knows none")

	fmt.Println(count)

}
