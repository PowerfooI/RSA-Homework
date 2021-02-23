package rsa

import "math/big"

var (
	p *big.Int
	q *big.Int
	n *big.Int
	phiN *big.Int
	e = big.NewInt(65537)
	d *big.Int
)

func debugSetup() {
	p, q = producePQ(768, 4)
	n = new(big.Int).Mul(p, q)
	phiN = new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))
	d, _, _ = extEuclid(new(big.Int).Set(e), new(big.Int).Set(phiN))
	if d.Cmp(zero) < 0 {
		d.Add(d, phiN)
	}
}