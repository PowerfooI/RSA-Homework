package rsa

import (
	"fmt"
	"testing"
)

func BenchmarkGenRSAKeys(b *testing.B) {
	SetUp()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenRSAKeys(768, 2)
	}
}

func TestPKCS1Padding(t *testing.T) {
	SetUp()
	pub, pri := GenRSAKeys(1024, 4)
	fmt.Printf("n = %v \np = %v \nq = %v \nd = %v \n", pub.NValue, pri.PValue, pri.QValue, pri.DValue)
	originalInput := "hello pig! 123456789012345678901234567890123456789012345678901234567890, what about the msg is of overlength? I mean it's not the solution for such a problem"

	blocks := PKCS1Padding(originalInput, 1024)
	ec := EncryptBlocks(pub, blocks)
	fmt.Println("og =", originalInput)
	fmt.Println("ec =", ec)
	deBlocks := DecryptBlocks(pri, ec)
	dc := PKCS1Depadding(deBlocks)
	fmt.Println("dc =", dc)
}

func Test2(t *testing.T) {
	count := 100
	for count > 0 {
		pub, pri := GenRSAKeys(256, 4)
		originalInput := "just for test"
		blocks := PKCS1Padding(originalInput, 256)
		ec := EncryptBlocks(pub, blocks)
		deBlocks := DecryptBlocks(pri, ec)
		fmt.Println(pri.PValue, pri.QValue)
		fmt.Println(blocks)
		fmt.Println(deBlocks)
		dc := PKCS1Depadding(deBlocks)
		if dc != originalInput {
			fmt.Println("Wrong result!")
		}
		count--
	}

}