package main

import (
	"fmt"
	"learn/utils"
	// "learn/utils"
	// "path/filepath"
	// "strings"
	// "os"
	// "io"
)

// var count = utils.WordCount("the one who knows one is one who knows none")


func main(){

	// sum := 0

	// f := utils.Fibonacci() 
	// for i := 0; sum < 4000000; i++ { 
	// 	num := f(i)
	// 	if num % 2 == 0 {
	// 		sum += num
	// 	}
	// }
	// fmt.Println(sum) 


	// s := strings.NewReader("V yvxr unz")
	// r := utils.NewRot13Reader(s)
	// io.Copy(os.Stdout, &r)

	// fmt.Println(utils.SumMul(3,5))

	// fmt.Println(utils.Primefactor(600851475143))
	// fmt.Println(utils.LgPalindrom())
	// fmt.Println(utils.FindDiffSquares(100))

	// dir, err := os.UserHomeDir()

	// if err != nil {
	// 	fmt.Println("could not get user home dir")
	// }
	
	// fmt.Println(utils.JumpDirectory("vendor", dir + "/Development/collaborations"))
	// var nums = []int{2,-1,1}

	// fmt.Println(utils.GetPrime(10001))

	// ".*?" ==> To(Equal(false)) ✔️
	// "a" ==> To(Equal(true)) ✔️
	// "Mazinkaiser" ==> To(Equal(true)) ✔️
	// "hello world_" ==> To(Equal(false)) ✔️
	// "     " ==> To(Equal(false)) ✔️
	// "PassW0rd" ==> To(Equal(true)) ✔️
	// "" ==> To(Equal(false))
	// "\n\t\n" ==> To(Equal(false))
	// "ciao\n$$_" ==> To(Equal(false))
	// "__ * __" ==> To(Equal(false))
	// "&)))(((" ==> To(Equal(false))
	// "43534h56jmTHHF3k" ==> To(Equal(true))

	fmt.Println(utils.Alphanumeric(""))

	


}
