package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	fmt.Println("Prime numbers in between 1 and 10 are:")
	for i := 1; i <= 10; i++ {
		if isPrime(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Printf("Sum of Prime numbers in between 1 and 10: %d\n", sum)
}
