package rsa

import "testing"

func BenchmarkGenRSAKeys(b *testing.B) {
	SetUp()
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		GenRSAKeys(768, 2)
	}
}