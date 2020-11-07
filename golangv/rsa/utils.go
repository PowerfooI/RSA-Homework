package rsa

import (
	"math/big"
)

func QuickExpMod(base big.Int, exponent big.Int, modulo big.Int) *big.Int {
	res := big.NewInt(1)
	for exponent.Cmp(big.NewInt(1))>0 {
		q, m := exponent.DivMod(&exponent, big.NewInt(2), nil)
		if m.Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, base.Mod(&base, &modulo))
		}
		base.Mul(&base, base.Mod(&base, &modulo))
		exponent = *q
	}
	return res
}

func ExtEuclid(a big.Int, b big.Int) (*big.Int, *big.Int, *big.Int) {
	if b.Cmp(big.NewInt(0)) == 0{
		return big.NewInt(1), big.NewInt(0), &a
	} else {
		u, v, q := ExtEuclid(b, *a.Mod(&a, &b))
		u, v = v, u.Sub(u, a.Div(&a, &b).Mul(&a, v))
		return u, v, q
	}
}