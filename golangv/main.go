package main

import (
	"fmt"
	"golangv/rsa"
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
	msg := "hello"
	fmt.Println("plain msg =", msg)
	c := rsa.EncodeMsg(pub, msg)
	fmt.Println("encrypt msg =", c)
	md := rsa.DecodeMsg(pri, c)
	fmt.Println("decrypt msg =", md)

}
