package rsa

import "math/big"

func GenRSAKeys(nBitSize, nGoRoutines int) (*PublicKey, *PrivateKey) {
	p, q := producePQ(nBitSize, nGoRoutines)
	n := new(big.Int).Mul(p, q)
	phiN := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))
	d, _, _ := extEuclid(new(big.Int).Set(e), new(big.Int).Set(phiN))
	if d.Cmp(zero) < 0 {
		d.Add(d, phiN)
	}
	pubKey := &PublicKey{
		NLength: nBitSize,
		NValue: n,
		EValue: big.NewInt(65537),
	}
	priKey := &PrivateKey{
		NLength: nBitSize,
		NValue: n,
		DValue: d,
		PValue: p,
		QValue: q,
	}
	return pubKey, priKey
}