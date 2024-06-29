package utils

import "fmt"


func SumMul(x , y int) int {
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
	return sum
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

func Primefactor(num int) int {
	result := num

	i := 2
	
	for i < result {
		for result % i == 0 {
			result = result / i
		}
		i = i + 1
	}

	return result
}


func isPalindrome(number int) bool {
    var reversed int
    var currentNumber = number

    for currentNumber > 0 {
        reversed = reversed*10 + currentNumber%10
        currentNumber /= 10
    }
    
    return reversed == number 
}

func LgPalindrom() int {
    var maxProduct int
    for i := 990; i >= 100; i -= 11 {
        for j := 999; j >= 100; j-- {
            product := i * j

            if product > maxProduct && isPalindrome(product) {
                maxProduct = product
				fmt.Println(i, j)
                break
            } else if maxProduct > product {
                break
            }
        }
    }
    return maxProduct
}
