package rsa

import "math/big"

type PublicKey struct {
	NLength int      `json:"n_len"`
	NValue  *big.Int `json:"n_val"`
	EValue  *big.Int `json:"e_val"`
}

type PrivateKey struct {
	NLength int      `json:"n_len"`
	NValue  *big.Int `json:"n_val"`
	DValue  *big.Int `json:"d_val"`
	PValue  *big.Int `json:"p_val"`
	QValue  *big.Int `json:"q_val"`
}

type Blocks struct {
	byteMatrix [][]byte
}

type Key interface {
	encodeBigInt(*big.Int) *big.Int
	getNLength() int
	IsPublic() bool
}

func (k *PublicKey) encodeBigInt(msg *big.Int) *big.Int {
	return quickExpMod(msg, k.EValue, k.NValue)
}

func (k *PublicKey) getNLength() int {
	return k.NLength
}

func (k *PublicKey) IsPublic() bool {
	return true
}

func (k *PrivateKey) encodeBigInt(msg *big.Int) *big.Int {
	return quickExpMod(msg, k.DValue, k.NValue)
}

func (k *PrivateKey) getNLength() int {
	return k.NLength
}

func (k *PrivateKey) IsPublic() bool {
	return false
}
