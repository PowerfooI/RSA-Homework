package main

import (
	"fmt"
	"golangv/rsa"
	"math/big"
	"time"
)

func main() {
	rsa.SetUp()

	start := time.Now()
	pub, pri :=	rsa.GenRSAKeys(1024, 4)
	//_, _ = rsa.GenRSAKeys(1024, 4)
	end := time.Now()
	fmt.Printf("total time = %.2fs\n", end.Sub(start).Seconds())


	// test for encrypt & decrypt
	//fmt.Printf("n = %v \np = %v \nq = %v \nd = %v \n", pub.NValue.Text(16), pri.PValue.Text(16), pri.QValue.Text(16), pri.DValue.Text(16))
	msg := big.NewInt(235345)
	fmt.Println("plain msg =", msg)
	c := rsa.EncryptMsg(pub, msg)
	fmt.Println("encrypt msg =", c)
	md := rsa.DecryptMsg(pri, c)
	fmt.Println("decrypt msg =", md)

	if md.Cmp(msg) != 0 {
		fmt.Println("Wrong results!")
	}
}
