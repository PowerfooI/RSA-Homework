package rsa

import (
	"math/big"
	"math/rand"
)

func GenRandomOddNum(pBitSize int) *big.Int {
	var res []uint8
	// 保证最高位是1
	res = append(res, '1')
	for i:=0;i<pBitSize-2;i++ {
		res = append(res, randSrc[rand.Intn(2)])
	}
	// 保证最低位是1，确保是奇数
	res = append(res, '1')
	ret := new(big.Int)
	ret.SetString(string(res), 2)
	return ret
}

func genRandomAinTest(targetNum *big.Int) *big.Int {
	// 保证比num更小
	nBits := targetNum.BitLen() - 1
	var res []uint8
	for i:=0;i<nBits;i++ {
		res = append(res, randSrc[rand.Intn(2)])
	}
	ret := new(big.Int)
	ret.SetString(string(res), 2)
	if ret.Cmp(one) <= 0 {
		ret.Add(ret, two)
	}
	return ret
}

func MillerRabinTest(p *big.Int, pBitSize int) bool {
	var times int
	switch {
	case pBitSize >= 1024:
		times = 3
	case pBitSize >= 512:
		times = 4
	case pBitSize >= 384:
		times = 5
	case pBitSize >= 256:
		times = 7
	default:
		times = 10
	}
	nMinusOne := new(big.Int).Sub(p, one)
	s := 0
	for (new(big.Int).And(nMinusOne,two)).Cmp(zero) == 0 {
		s++
		nMinusOne.Div(nMinusOne, two)
	}
	d := nMinusOne

	pMinusOne := new(big.Int).Sub(p, one) // p - 1
	for times > 0 {
		a := genRandomAinTest(new(big.Int).Sub(p, one))
		r1 := QuickExpMod(new(big.Int).Set(a), new(big.Int).Set(d), new(big.Int).Set(p))
		r1isOdd := r1.Cmp(one) == 0
		rCount := 0
		for r:= 0;r<s;r++ {
			r2 := QuickExpMod(new(big.Int).Set(a), new(big.Int).Mul(QuickExp(big.NewInt(2), big.NewInt(int64(r))), d), new(big.Int).Set(p))
			if !r1isOdd && r2.Cmp(pMinusOne) != 0 {
				rCount++
			} else {
				break
			}
		}
		if rCount == s {
			return false
		}
		times--
	}
	return true
}

func GenRandomPrimeNumber(pBitSize int) *big.Int {
	p := GenRandomOddNum(pBitSize)
	for {
		primeFlag := true
		for _, tp := range primeNumUnder11k {
			if (new(big.Int).Mod(p, tp)).Cmp(zero) == 0 {
				primeFlag = false
				break
			}
		}
		if !primeFlag {
			p.Add(p, two)
			continue
		}
		if !MillerRabinTest(p, pBitSize) {
			p.Add(p, two)
			continue
		} else {
			return p
		}
	}
}

func ProducePQ(nBitSize, nGoroutines int) (*big.Int, *big.Int) {
	kBits := nBitSize / 2
	pqChan := make(chan *big.Int)
	for i:=0;i<nGoroutines;i++ {
		go func() {
			for {
				select {
					case pqChan <- GenRandomPrimeNumber(kBits):
						continue
				}
			}
		}()
	}
	var firstPrime *big.Int
	for {
		select {
		case primeNum := <-pqChan:
			if firstPrime != nil {
				if primeNum.Cmp(firstPrime) != 0 {
					return firstPrime, primeNum
				}
			} else {
				firstPrime = primeNum
			}
		}
	}
}