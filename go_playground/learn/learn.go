package main

import (
	"learn/utils"
	"fmt"
	"strings"
	"os"
	"io"
)

var count = utils.WordCount("the one who knows one is one who knows none")


func main(){

	f := utils.Fibonacci() 
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
	fmt.Println(count) 

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := utils.NewRot13Reader(s)
	io.Copy(os.Stdout, &r)



}