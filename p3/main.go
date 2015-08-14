package main

func main() {
	n := 600851475143
	d := 2
	for d <= n {
		if n%d == 0 {
			n = n / d
		} else {
			if d > 2 {
				d += 2
			} else {
				d = 3
			}
		}
	}
	println(d)
}
