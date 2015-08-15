package main

import (
	"fmt"
	"math/big"
)

func main() {
	k := 20

	i := big.NewInt(0)
	fmt.Println(i.Binomial((int64)(k*2), (int64)(k)))
}
