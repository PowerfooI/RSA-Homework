package rsa

import (
	"math/big"
)

var randSrc = "01"

func QuickExp(base, exponent *big.Int) *big.Int {
	res := big.NewInt(1)
	for exponent.Cmp(big.NewInt(1))>0 {
		q, m := exponent.DivMod(exponent, big.NewInt(2), new(big.Int))
		if m.Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, base)
		}
		base.Mul(base, base)
		exponent = q
	}
	return res
}

func QuickExpMod(base *big.Int, exponent *big.Int, modulo *big.Int) *big.Int {
	res := big.NewInt(1)
	for exponent.Cmp(big.NewInt(1))>0 {
		q, m := exponent.DivMod(exponent, big.NewInt(2), new(big.Int))
		if m.Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, new(big.Int).Mod(base, modulo))
		}
		base.Mul(base, new(big.Int).Mod(base, modulo))
		exponent = q
	}
	return res
}

func ExtEuclid(a *big.Int, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if b.Cmp(big.NewInt(0)) == 0{
		return big.NewInt(1), big.NewInt(0), a
	} else {
		u, v, q := ExtEuclid(b, new(big.Int).Mod(a, b))
		u, v = v, new(big.Int).Sub(u, new(big.Int).Mul(new(big.Int).Div(a, b), v))
		return u, v, q
	}
}

