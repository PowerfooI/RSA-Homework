package rsa

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

func encodeBlocks(key Key, blocks *Blocks) string {
	res := ""
	for _, b := range blocks.byteMatrix {
		tmpNum := new(big.Int).SetBytes(b)
		ec := key.encodeBigInt(tmpNum)
		blockOutput := hex.EncodeToString(ec.Bytes())
		res += blockOutput
	}
	return res
}

func decodeBlocks(key Key, deText string) *Blocks {
	blockByteLength := key.getNLength() / 8
	dc, _ := hex.DecodeString(deText)

	lenOfInput := len(dc)
	dcb := &Blocks{}
	byteMatrix := make([][]byte, 0)

	tmpSize := blockByteLength
	for cur:=0;cur<lenOfInput; {
		if cur + blockByteLength > lenOfInput {
			tmpSize = lenOfInput - cur
		}
		encryptBlock := dc[cur:cur+tmpSize]
		tmpMsg := new(big.Int)
		tmpMsg.SetBytes(encryptBlock)
		tmpMsg = key.encodeBigInt(tmpMsg)

		msgBitLen := tmpMsg.BitLen()
		zeroFlag := msgBitLen % 8 == 0
		byteLen := msgBitLen / 8
		if !zeroFlag {
			byteLen++
		}
		gap := blockByteLength - byteLen
		var frontZeros []byte
		for i:=0;i<gap;i++ {
			frontZeros = append(frontZeros, byte(0))
		}
		byteMatrix = append(byteMatrix, append(frontZeros, tmpMsg.Bytes()...))
		cur += tmpSize
	}

	dcb.byteMatrix = byteMatrix
	return dcb
}

func EncodeMsg(key Key, input string) string {
	blocks := paddingPKCS1(input, key.getNLength(), key.isPublic())
	encodedStr := encodeBlocks(key, blocks)
	return encodedStr
}

func DecodeMsg(key Key, input string) string {
	blocks := decodeBlocks(key, input)
	decodedStr := depaddingPKCS1(blocks)
	return decodedStr
}

func Sign(key Key, input string) string {
	m := md5.New()
	m.Write([]byte(input))
	digest := m.Sum(nil)
	signature := EncodeMsg(key, string(digest))
	return signature
}

