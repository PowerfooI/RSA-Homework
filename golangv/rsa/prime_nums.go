package rsa

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

var primeNumUnder11k []*big.Int

func SetUp() {
	rand.Seed(time.Now().Unix())
	genPreTestList()
}

func genPreTestList() {
	for i:=2;i<1.1e4;i++ {
		primeFlag := true
		for j:=i;j<int(math.Sqrt(float64(i)))+1;j++ {
			if i % j == 0 {
				primeFlag = false
				break
			}
		}
		if primeFlag {
			primeNumUnder11k = append(primeNumUnder11k, big.NewInt(int64(i)))
		}
	}
}

func GenRandomOddNum(nBits int) *big.Int {
	var res []uint8
	// 保证最高位是1
	res = append(res, '1')
	for i:=0;i<nBits-2;i++ {
		res = append(res, randSrc[rand.Intn(2)])
	}
	// 保证最低位是1，确保是奇数
	res = append(res, '1')
	ret := new(big.Int)
	ret.SetString(string(res), 2)
	return ret
}

func genRandomAinTest(num *big.Int) *big.Int {
	// 保证比num更小
	nBits := num.BitLen() - 1
	var res []uint8
	for i:=0;i<nBits-1;i++ {
		res = append(res, randSrc[rand.Intn(2)])
	}
	res = append(res, '1')
	ret := new(big.Int)
	ret.SetString(string(res), 2)
	return ret
}

func millerRabinTest(s int, p, d *big.Int, nBitSize int) bool {
	var times int
	switch {
	case nBitSize >= 1024:
		times = 3
	case nBitSize >= 512:
		times = 4
	case nBitSize >= 384:
		times = 5
	case nBitSize >= 256:
		times = 7
	default:
		times = 10
	}
	for times > 0 {
		a := genRandomAinTest(new(big.Int).Sub(p, big.NewInt(1)))
		for r:= 0;r<s;r++ {
			r1 := QuickExpMod(a, d, p)
			r2 := QuickExpMod(a, new(big.Int).Mul(QuickExp(big.NewInt(2), big.NewInt(int64(r))), d), p)
			if r1.Cmp(big.NewInt(1)) != 0 && r2.Cmp(big.NewInt(1)) != 0 {
				return false
			}
		}
		times--
	}
	return true
}

func GenRandomPrimeNumber(nBits int) *big.Int {
	p := GenRandomOddNum(nBits)
	for {
		primeFlag := true
		for _, tp := range primeNumUnder11k {
			if (new(big.Int).Mod(p, tp)).Cmp(big.NewInt(0)) == 0 {
				primeFlag = false
				break
			}
		}
		if !primeFlag {
			p.Add(p, big.NewInt(2))
			continue
		}
		nMinusOne := new(big.Int).Sub(p, big.NewInt(1))
		s := 0
		for (new(big.Int).Mod(nMinusOne, big.NewInt(2))).Cmp(big.NewInt(0)) == 0 {
			s++
			nMinusOne.Div(nMinusOne, big.NewInt(2))
		}
		d := nMinusOne
		if !millerRabinTest(s, p, d, nBits) {
			p.Add(p, big.NewInt(2))
			continue
		} else {
			return p
		}
	}
}

func ProducePQ(nBitSize, nGoroutines int) (*big.Int, *big.Int) {
	nBits := nBitSize / 2
	pqChan := make(chan *big.Int)
	doneChan := make(chan bool)
	for i:=0;i<nGoroutines;i++ {
		go func() {
			for {
				select {
					case <-doneChan:
						return
					case pqChan <- GenRandomPrimeNumber(nBits):
						fmt.Println("Produce one big prime number!")
				}
			}
		}()
	}
	var pqArray []*big.Int
	pqArray = make([]*big.Int, 0)
	for {
		select {
		case primeNum := <-pqChan:
			if len(pqArray) > 0 {
				if primeNum.Cmp(pqArray[0]) != 0 {
					pqArray = append(pqArray, primeNum)
					for i:=0;i<nGoroutines;i++ {
						doneChan <- true
					}
					return pqArray[0], pqArray[1]
				}
			} else {
				pqArray = append(pqArray, primeNum)
			}
		}
	}
}