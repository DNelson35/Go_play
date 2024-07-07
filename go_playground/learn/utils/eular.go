package utils

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	// "time"
	// "sort"
)


func SumMul(x , y int) int {
	sum := 0
	for i:=0; i < 1000; i++ {
		if i % x == 0 || i % y == 0 {
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

	i := 3
	
	for i < result {
		if result % i == 0 {
			result = result / i
		}
		i = i + 2
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


// 1 2 3 2^2 5 3 2 7 2^3 3^2 5 2 11 2^2 3 13 7 2 5 3 2^4 17 2 3^2 19 2^2 5 all prime factors
// 1 2^4 3^2 5 7 11 13 17 19 greatest prime factors

// project euler answer number 5 232792560

func FindDiffSquares(x int) int{
	// start := time.Now()
	sumsq := 0
	sqsum := 0

	for i:=0; i <= x; i++ {
		sumsq += int(math.Pow(float64(i), 2))
		sqsum += i
	}
	sqsum = int(math.Pow(float64(sqsum), 2))

	// fmt.Println(time.Since(start).Nanoseconds())
	return sqsum - sumsq
}

// current problem: one i need to check to see if it is a system directory before trying to access i belive that is one part of the problem as well as it is being appended to the dirstosearch which i dont want. second problem I need to rethink how i stor the dirctories to seach or how to get the path to the dir that holds them i belive the problem is that im not currently in the directory and im tring to read form a directory nested so it cant see the direcotry im tring to acess



func JumpDirectory(name string) string {
	userdir, err := os.UserHomeDir()
	dirstoSearch := [][]fs.DirEntry{}

	if err != nil{
		fmt.Println("error getting user directory: ", err)
	}
	
	os.Chdir(userdir)
	cleanDirs, path := getDirs(userdir, name)

	if path != ""{
		return path
	}

	path = FileTree(name, cleanDirs, &dirstoSearch)

	for _, dirset := range dirstoSearch {
		FileTree(name, dirset, &dirstoSearch)
	}
	// fmt.Println(dirstoSearch)
	return path
}

func FileTree(name string, cleanDirs []fs.DirEntry, dirstoSearch *[][]fs.DirEntry)string{
	fmt.Println(cleanDirs)
	for _, dir := range cleanDirs{
		// os.Chdir(dir.Name())
		fmt.Println(filepath.Abs(dir.Name()))
		dirs, path := getDirs(dir.Name(), name)
		if path != ""{
			// fmt.Println(path)
			return path
		}
		*dirstoSearch = append(*dirstoSearch, dirs)
	}

	// if len(*dirstoSearch) > 0 {
	// 	*dirstoSearch = (*dirstoSearch)[1:]
	// }
	
	return ""
}

func searchVisDirs(dirs []fs.DirEntry, name string)([]fs.DirEntry, string) {
	var visDirs []fs.DirEntry
	for _, dir := range dirs{
		if dir.Name()[0] == '.' || !dir.IsDir(){
			continue
		}else if checkMatch(name, dir){
			path, _ := filepath.Abs(dir.Name())
			return nil, path
		}

		if len(dirs) != 0{
			visDirs = append(visDirs, dir)
		}
		
	}
	return visDirs, ""
}

func checkMatch(name string, dir fs.DirEntry)bool{		
	return dir.Name() == name 
	
}

func getDirs(dir string, name string)([]fs.DirEntry, string){
	dirs, err := os.ReadDir(dir)

	// fmt.Println(dirs)
	
	if err != nil {
		fmt.Println(err)
	}

	return searchVisDirs(dirs, name)
}




