package rsa

import "math/big"

type PublicKey struct {
	Length int
	NValue *big.Int
	EValue *big.Int
}

type PrivateKey struct {
	Length int
	NValue *big.Int
	DValue *big.Int
	PValue *big.Int
	QValue *big.Int
}

type Blocks struct {
	byteMatrix  [][]byte
}
