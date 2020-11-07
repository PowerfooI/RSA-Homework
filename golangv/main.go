package main

import (
	"fmt"
	"golangv/rsa"
	"math/big"
)

func main() {
	rsa.SetUp()
	fmt.Println(rsa.ExtEuclid(big.NewInt(189), big.NewInt(321)))
	fmt.Println(rsa.ProducePQ(128, 5))
}