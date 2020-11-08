package rsa

import "math/big"

func EncryptMsg(key *PublicKey, msg *big.Int) *big.Int {
	return QuickExpMod(new(big.Int).Set(msg), new(big.Int).Set(key.EValue), new(big.Int).Set(key.NValue))
}

func DecryptMsg(key *PrivateKey, eMsg *big.Int) *big.Int {
	return QuickExpMod(new(big.Int).Set(eMsg), new(big.Int).Set(key.DValue), new(big.Int).Set(key.NValue))
}

func EncryptBlocks(key *PublicKey, blocks *Blocks) string {
	res := ""
	blockByteLength := key.Length / 4
	blockOutputLength := blockByteLength * 2
	for _, b := range blocks.byteMatrix {
		tmpNum := new(big.Int).SetBytes(b)
		ec := EncryptMsg(key, tmpNum)
		blockOutput := ec.Text(16)
		outputLen := len(blockOutput)
		if outputLen < blockOutputLength {
			var tmpBytes string
			for i:=outputLen;i<blockOutputLength;i++ {
				tmpBytes += "0"
			}
			blockOutput = tmpBytes + blockOutput
		}
		res += blockOutput
	}
	return res
}

func DecryptBlocks(key *PrivateKey, dc string) *Blocks {
	blockByteLength := key.Length / 4
	blockOutputLength := blockByteLength * 2

	lenOfInput := len(dc)
	dcb := &Blocks{}
	byteMatrix := make([][]byte, 0)

	tmpSize := blockOutputLength
	for cur:=0;cur<lenOfInput; {
		if cur + blockOutputLength > lenOfInput {
			tmpSize = lenOfInput - cur
		}
		encryptBlock := dc[cur:cur+tmpSize]
		tmpMsg := new(big.Int)
		tmpMsg.SetString(encryptBlock, 16)

		tmpMsg = DecryptMsg(key, tmpMsg)
		byteMatrix = append(byteMatrix, tmpMsg.Bytes())

		cur += tmpSize
	}

	//rawHexInput := new(big.Int)
	//rawHexInput.SetString(dc, 16)
	//dcBytes := rawHexInput.Bytes()
	//
	//blockSize := key.Length / 8

	dcb.byteMatrix = byteMatrix
	return dcb
}