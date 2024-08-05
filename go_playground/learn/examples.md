# Examples

## learn/utils/euler.go

## SumMul
takes in two integers and returns the sum of all multiples of the integers under 1000. in the example we pass 3 and 5 the function will initiates a loop up to 1000 and use the modulus operator to determine if the index of the current iteration has a remainder if divided by either number if the number divides evenly the index is then added to the sum when the loop finishes the sum is returned.

```go
utils.SumMul(3,5) --> 233168
```

### Fibonacci

in this example we calculate the sum of all even fibonacci numbers thats value do not exceed 4000000. Fibonacci returns a function that takes in a integer from an sequence and returns the current fibonacci number in that sequence. In the example we call Fibonacci to assign its return function to a variable we then loop from 0 to the nth fibonacci number we wish to obtain. On each call the fibonacci number will advance to the next number in the sequence the number returned will then be checked to determine if the number is even if it is we add it to the sum if not we continue effectively getting the sum of all fibonacci numbers up to the nth fibonacci number in the sequence.

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
### PrimeFactor

PrimeFactor in this example takes in the integer 2062 and returns the largest prime factor which is 1031. In this case 2062 has two prime factors 2 and 1031, and PrimeFactor returns the largest prime factor.

```go
utils.Primefactor(2062) --> 1031
```

## learn/utils/rot13.go

### NewRot13Reader - Read

This example showcases the utility of a Rot13Reader from the utils package, which seamlessly integrates ROT13 encoding into data processing workflows. By wrapping around either a file or standard input (stdin), the Rot13Reader applies ROT13 encoding on-the-fly as data is read. This flexibility allows for dynamic encoding of input data in real-time, whether it's streaming from stdin or loaded from a file. The example demonstrates reading and printing encoded lines of text from stdin, highlighting the ease of implementing ROT13 encoding in various contexts.

```go
reader := utils.NewRot13Reader(os.Stdin)

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic("read error")
			}
			break
		}
		line := string(buf[:n])
		fmt.Println(line)
	}
// input: this is an example --> Guvf vf na rknzcyr
```

