package utils

import (
	"fmt"
	"math"
)

// SumMul calculates the sum of all integers between 0 and 999 (inclusive)
// that are divisible by either x or y. It iterates through each integer in the range,
// checks if the current integer is divisible by both x and y using the modulus operator,
// and if so, adds it to the running total stored in 'sum'. Finally, it returns the
// accumulated sum after completing the loop.
func SumMul(x , y int) int {
	sum := 0
	for i:=0; i < 1000; i++ {
		if i % x == 0 || i % y == 0 {
			sum += i
		}
	}
	return sum
}

// Fibonacci generates a closure that calculates the nth Fibonacci number when invoked.
// This function initializes the first two positions (pos1 and pos2) of the Fibonacci sequence,
// where pos1 represents the current position and pos2 represents the next position in the sequence.
// Each call to the returned function advances the sequence one position forward,
// starting from the second position (index 1), and thus requires iteration to reach the desired index.
// Specifically, the sequence begins with 0 and 1, and subsequent numbers are the sum of the preceding two.
// Therefore, to obtain the actual Fibonacci number at a given index,
// the returned function must be called within a loop iterating from the initial index up to and including the desired index.
// Note: The function does not directly return the nth Fibonacci number based on the input index alone;
// instead, it relies on sequential calls to advance through the sequence.
func Fibonacci() func(x int) int { 
	pos1, pos2 := 0,1

	return func(x int) int {
		if x == 1 {
			return pos1 + pos2
		}else if x < 1 {
			return 0
		}
		num := pos1 + pos2
		pos1, pos2 = pos2, num
		return num
	}
}

// Primefactor calculates the largest prime factor of a given positive integer.
//
// The function initializes the result with the input number and iterates from 3 upwards,
// incrementing by 2 each time to skip even numbers (since even numbers greater than 2 cannot be prime).
// For each iteration, if the result is divisible by the current iterator (i), it divides the result by i.
// If the result of the division is divisible by two, the result is changed to be equal to i, capturing the prime factor.
// This process continues until the iterator exceeds the result, at which point the result contains the largest prime factor.
func Primefactor(num int) int {
	result := num

	i := 3
	
	for i < result {
		if result % i == 0 {
			result = result / i
			if result % 2 == 0{
				result = i
			}
		}
		i = i + 2
	}

	return result
}

// Helper function to check if a number is a palindrome
func isPalindrome(number int) bool {
    var reversed int
    var currentNumber = number

    for currentNumber > 0 {
        reversed = reversed*10 + currentNumber%10
        currentNumber /= 10
    }
    
    return reversed == number 
}

// Function to find the largest palindrome product of two 3-digit numbers
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


// 1 2 3 2^2 5 3 2 7 2^3 3^2 5 2 11 2^2 3 13 7 2 5 3 2^4 17 2 3^2 19 2^2 5 all prime factors
// 1 2^4 3^2 5 7 11 13 17 19 greatest prime factors

// project euler answer number 5 232792560


// This function calculates the difference between the squares of the first x numbers and the square of the sum of the first x numbers. argument x is the bounds for number you would like this calculation to go up to.
func FindDiffSquares(x int) int{
	sumsq := 0
	sqsum := 0

	for i:=0; i <= x; i++ {
		sumsq += int(math.Pow(float64(i), 2))
		sqsum += i
	}
	sqsum = int(math.Pow(float64(sqsum), 2))

	return sqsum - sumsq
}


func checkPrime(num int) bool{
	if num <= 1 {return false}
	if num <= 3 {return true}
	if num % 2 == 0 || num % 3 == 0 {return false}
	for i:=5; i * i <= num; i+= 6{
		if num % i == 0|| num % (i+2) == 0 {return false}
	}
	return true
}

func GetPrime(n int) int{
    i := 0
    num := 1
    for i < n {
        num++
        if (checkPrime(num)) {
            i++
        }
    }
    return num
}






