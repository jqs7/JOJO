package main

import (
	"math/big"
	"strconv"
)

func main() {
	k := 1000
	sum := 0
	i := big.NewInt(2)
	result := i.Exp(i, big.NewInt(int64(k)), nil).String()
	for k := range result {
		i, _ := strconv.Atoi(string(result[k]))
		sum += i
	}
	println(sum)
}
