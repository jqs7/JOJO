package main

func main() {
	k := 1000
	for a := 1; a < k/3+1; a++ {
		for b := 1; b < k/2+1; b++ {
			c := k - a - b
			if a*a+b*b == c*c {
				println(a * b * c)
			}
		}
	}
}
