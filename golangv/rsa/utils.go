package rsa

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

var randSrc = "01"

var primeNumUnder11k []*big.Int

// used for non-side effect ops
var (
	zero = big.NewInt(0)
	one = big.NewInt(1)
	two = big.NewInt(2)
)

func SetUp()  {
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

func QuickExp(base, exponent *big.Int) *big.Int {
	res := big.NewInt(1)
	for exponent.Cmp(one)>0 {
		q, m := exponent.DivMod(exponent, two, new(big.Int))
		if m.Cmp(one) == 0 {
			res.Mul(res, base)
		}
		base.Mul(base, base)
		exponent = q
	}
	return res
}

func QuickExpMod(base, exponent, modulo *big.Int) *big.Int {
	res := big.NewInt(1)
	for exponent.Cmp(zero)>0 {
		q, m := exponent.DivMod(exponent, two, new(big.Int))
		if m.Cmp(one) == 0 {
			res.Mul(res, base).Mod(res, modulo)
		}
		base.Mul(base, base).Mod(base, modulo)
		exponent = q
	}
	return res
}

func ExtEuclid(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if b.Cmp(zero) == 0{
		return big.NewInt(1), big.NewInt(0), new(big.Int).Set(a)
	} else {
		u, v, q := ExtEuclid(b, new(big.Int).Mod(a, b))
		u, v = v, new(big.Int).Sub(u, new(big.Int).Mul(new(big.Int).Div(a, b), v))
		return u, v, q
	}
}

