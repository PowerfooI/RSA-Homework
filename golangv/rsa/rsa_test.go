package rsa

import (
	"fmt"
	"testing"
)

func BenchmarkGenRSAKeys(b *testing.B) {
	SetUp()
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		GenRSAKeys(768, 2)
	}
}

func TestPKCS1Padding(t *testing.T) {
	SetUp()
	pub, pri := GenRSAKeys(768, 4)
	fmt.Printf("n = %v \np = %v \nq = %v \nd = %v \n", pub.NValue, pri.PValue, pri.QValue, pri.DValue)
	originalInput := "hello pig! 123456789012345678901234567890123456789012345678901234567890"

	blocks := PKCS1Padding(originalInput, 512)
	ec := EncryptBlocks(pub, blocks)
	fmt.Println("og =", originalInput)
	fmt.Println("ec =", ec)
	deBlocks := DecryptBlocks(pri, ec)
	dc := PKCS1Depadding(deBlocks, 512)
	fmt.Println("dc =", dc)
}

func Test1(t *testing.T) {
	fmt.Println([]byte{'1', '2'})
}