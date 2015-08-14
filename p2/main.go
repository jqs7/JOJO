package main

func main() {
	f1 := 0
	f2 := 1
	sum := 0
	for {
		f1, f2 = f2, f1+f2
		if f2 > 4000000 {
			break
		}
		if f2%2 == 0 {
			sum += f2
		}
	}
	println(sum)
}
