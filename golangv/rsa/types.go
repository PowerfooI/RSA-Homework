package rsa

import "math/big"

type PublicKey struct {
	NLength int
	NValue  *big.Int
	EValue  *big.Int
}

type PrivateKey struct {
	NLength int
	NValue  *big.Int
	DValue  *big.Int
	PValue  *big.Int
	QValue  *big.Int
}

type Blocks struct {
	byteMatrix [][]byte
}

type Key interface {
	encodeBigInt(*big.Int) *big.Int
	getNLength() int
	isPublic() bool
}


func (k *PublicKey) encodeBigInt(msg *big.Int) *big.Int {
	return quickExpMod(msg, k.EValue, k.NValue)
}

func (k *PublicKey) getNLength() int {
	return k.NLength
}

func (k *PublicKey) isPublic() bool {
	return true
}

func (k *PrivateKey) encodeBigInt(msg *big.Int) *big.Int {
	return quickExpMod(msg, k.DValue, k.NValue)
}

func (k *PrivateKey) getNLength() int {
	return k.NLength
}

func (k *PrivateKey) isPublic() bool {
	return false
}

