package main

import (
	"learn/utils"
	"fmt"
)

var count = utils.WordCount("the one who knows one is one who knows none")


func main(){

	f := utils.Fibonacci() 
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
	fmt.Println(count) 

}