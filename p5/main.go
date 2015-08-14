package main

func divid(n int) bool {
	for j := 11; j <= 20; j++ {
		if n%j != 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 2520
	for {
		if divid(i) {
			break
		}
		i++
	}
	println(i)
}
