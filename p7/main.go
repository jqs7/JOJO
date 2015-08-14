package main

func main() {
	i, j, n := 2, 3, 10001
	for {
		j = j + 1
		if j > i/j {
			n--
			if n == 0 {
				break
			}
			j = 1
		}
		if i%j == 0 {
			i++
			j = 1
		}
	}
	println(i)
}
