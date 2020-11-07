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

