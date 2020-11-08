package rsa

import "math/big"

func EncryptMsg(key *PublicKey, msg *big.Int) *big.Int {
	return QuickExpMod(new(big.Int).Set(msg), new(big.Int).Set(key.EValue), new(big.Int).Set(key.NValue))
}

func DecryptMsg(key *PrivateKey, eMsg *big.Int) *big.Int {
	return QuickExpMod(new(big.Int).Set(eMsg), new(big.Int).Set(key.DValue), new(big.Int).Set(key.NValue))
}

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