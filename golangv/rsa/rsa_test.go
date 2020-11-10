package rsa

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestPKCS1Padding(t *testing.T) {
	SetUp()
	pub, pri := GenRSAKeys(512, 4)
	fmt.Printf("n = %v \np = %v \nq = %v \nd = %v \n", pub.NValue, pri.PValue, pri.QValue, pri.DValue)
	originalInput := "hello pig! 123456789012345678901234567890123456789012345678901234567890, what about the msg is of overlength? I mean it's not the solution for such a problem"

	blocks := paddingPKCS1(originalInput, 512, true)
	ec := encodeBlocks(pub, blocks)
	fmt.Println("og =", originalInput)
	fmt.Println("ec =", ec)
	deBlocks := decodeBlocks(pri, ec)
	dc := depaddingPKCS1(deBlocks)
	fmt.Println("dc =", dc)
}

func Test2(t *testing.T) {
	count := 10
	for count > 0 {
		pub, pri := GenRSAKeys(256, 4)
		originalInput := "12345678901234567890123456789012345678901234567890123"
		blocks := paddingPKCS1(originalInput, 256, true)
		ec := encodeBlocks(pub, blocks)
		deBlocks := decodeBlocks(pri, ec)
		fmt.Println(pri.PValue, pri.QValue)
		fmt.Println(blocks)
		fmt.Println(deBlocks)
		dc := depaddingPKCS1(deBlocks)
		if dc != originalInput {
			fmt.Println("Wrong result!")
		}
		count--
	}

}

func TestSignature(t *testing.T) {
	pub, pri := GenRSAKeys(1024, 4)
	originalInput := "hello digital signature"
	signature := Sign(pri, originalInput)
	fmt.Println(signature)
	m := md5.New()
	m.Write([]byte(originalInput))
	myDigest := string(m.Sum(nil))
	originalDigest := DecodeMsg(pub, signature)
	fmt.Printf("%x\n", myDigest)
	fmt.Printf("%x\n", originalDigest)
}
