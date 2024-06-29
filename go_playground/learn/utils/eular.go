package utils

import "fmt"

func SumMul(x , y int) {
	sum := 0
	for i:=0; i < 1000; i++ {
		if i % x == 0 && i % y == 0 {
			sum += i
			continue
		}else if i % x == 0 {
			sum += i
			continue
		}else if i % y == 0 {
			sum += i
			continue
		}
	}
	fmt.Println(sum)
}


func Fibonacci() func(x int) int { // for eular project find sum of all even numbers in the fib sequence that sum is less than 4000000 with a closure
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

func Primefactor(num int) {
	result := num

	i := 2
	
	for i < result {
		for result % i == 0 {
			result = result / i
		}
		i = i + 1
	}

	fmt.Println(result)
}

