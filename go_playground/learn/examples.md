# Examples

## learn/utils/euler

## SumMul
takes in two integers and returns the sum of all multiples of the integers under 1000. in the example we pass 3 and 5 the function will initiates a loop up to 1000 and use the modulus operator to determine if the index of the current iteration has a remainder if divided by either number if the number divides evenly the index is then added to the sum when the loop finishes the sum is returned.

```go
utils.SumMul(3,5) --> 233168
```

### Fibonacci

in this example we calculate the sum of all even fibonacci numbers thats value do not exceed 4000000
```go
sum := 0

	f := utils.Fibonacci() 
	for i := 0; sum < 4000000; i++ { 
    num := f(i)
	  if num % 2 == 0 {
			sum += num
		}
	}
	fmt.Println(sum) --> 4613732
```


