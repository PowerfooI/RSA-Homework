package rsa

import "math/big"

func GenRSAKeys(nBitSize, nGoRoutines int) (*PublicKey, *PrivateKey) {
	p, q := ProducePQ(nBitSize, nGoRoutines)
	n := new(big.Int).Mul(p, q)
	phiN := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))
	d, _, _ := ExtEuclid(new(big.Int).Set(e), new(big.Int).Set(phiN))
	if d.Cmp(zero) < 0 {
		d.Add(d, phiN)
	}
	pubKey := &PublicKey{
		Length: nBitSize / 2,
		NValue: n,
		EValue: big.NewInt(65537),
	}
	priKey := &PrivateKey{
		Length: nBitSize / 2,
		NValue: n,
		DValue: d,
		PValue: p,
		QValue: q,
	}
	return pubKey, priKey
}