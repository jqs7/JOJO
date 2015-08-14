package main

import "strconv"

func isPalindromic(n int) bool {
	ns := strconv.Itoa(n)
	for k := range ns {
		if k > len(ns)/2 {
			break
		}
		if ns[k] != ns[len(ns)-k-1] {
			return false
		}
	}
	return true
}

func main() {
	palindromic := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			k := i * j
			if isPalindromic(k) && k > palindromic {
				palindromic = k
			}
		}
	}
	println(palindromic)
}
