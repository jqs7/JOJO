package main

func main() {
	startNum, longest := 1, 0
	for i := 0; i < 1000000; i++ {
		count := chainLength(i)
		if count > longest {
			longest = count
			startNum = i
		}
	}
	println(startNum)
}

func chainLength(n int) int {
	count := 1
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	return count
}
