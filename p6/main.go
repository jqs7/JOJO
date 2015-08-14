package main

func main() {
	sumOfSq, sum := 0, 0
	for i := 1; i <= 100; i++ {
		sumOfSq += i * i
		sum += i
	}

	SqOfSum := sum * sum
	println(SqOfSum - sumOfSq)
}
