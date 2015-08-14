package main

import "math"

func main() {
	k := 500
	tnum := 0
	for i := 1; ; i++ {
		tnum += i
		sqrt := int(math.Sqrt(float64(tnum)))
		counter := 0
		for j := 1; j <= sqrt; j++ {
			if tnum%j == 0 {
				counter++
			}
		}
		if counter*2 > k {
			break
		}
	}
	println(tnum)
}
