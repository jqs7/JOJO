package main

import "math"

func main() {
	sum := 0
	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	println(sum)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	var sqrt, i int
	sqrt = int(math.Sqrt(float64(num)))
	for i = 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
