package rsa

import (
	"errors"
	"math/rand"
)

var (
	paddingSize = 11
)

func paddingPKCS1(rawTextInput string, nBitSize int, isPublic bool) *Blocks {
	nByteSize := nBitSize / 8
	pieceSize := nByteSize - paddingSize
	paddedBlocks := &Blocks{}

	input := []byte(rawTextInput)

	byteMatrix := make([][]byte, 0)
	lenOfInput := len(input)

	tmpSize := pieceSize
	for cur := 0; cur < lenOfInput; {
		var paddingContent []byte
		if isPublic {
			paddingContent = []byte{0x00, 0x02}
			// 11 - 2 - 1 = 8
			for i := 0; i < 8; i++ {
				paddingContent = append(paddingContent, byte(rand.Intn(0xff)))
			}
		} else {
			paddingContent = []byte{0x00, 0x01}
			// 11 - 2 - 1 = 8
			for i := 0; i < 8; i++ {
				paddingContent = append(paddingContent, byte(0xff))
			}
		}
		paddingContent = append(paddingContent, byte(0x00))

		if cur+pieceSize > lenOfInput {
			tmpSize = lenOfInput - cur
			newPaddingSize := pieceSize - tmpSize
			var backPaddingContent []byte
			for i := 0; i < newPaddingSize; i++ {
				backPaddingContent = append(backPaddingContent, byte(newPaddingSize))
			}
			tmpByteArray := make([]byte, tmpSize)
			copy(tmpByteArray, input[cur:cur+tmpSize])
			byteMatrix = append(byteMatrix, append(paddingContent, append(tmpByteArray, backPaddingContent...)...))
		} else if cur+pieceSize == lenOfInput {
			tmpByteArray := make([]byte, tmpSize)
			copy(tmpByteArray, input[cur:cur+tmpSize])
			byteMatrix = append(byteMatrix, append(paddingContent, tmpByteArray...))

			// 恰好相等时填充一整块
			var backPaddingContent []byte
			for i := 0; i < pieceSize; i++ {
				backPaddingContent = append(backPaddingContent, byte(pieceSize))
			}
			byteMatrix = append(byteMatrix, append(paddingContent, backPaddingContent...))
		} else {
			tmpByteArray := make([]byte, tmpSize)
			copy(tmpByteArray, input[cur:cur+tmpSize])
			byteMatrix = append(byteMatrix, append(paddingContent, tmpByteArray...))
		}
		cur += tmpSize
	}
	paddedBlocks.byteMatrix = byteMatrix
	return paddedBlocks
}

func depaddingPKCS1(blocks *Blocks) (string, error) {
	nBlocks := len(blocks.byteMatrix)
	blockLen := len(blocks.byteMatrix[0])
	var res []byte
	for i, b := range blocks.byteMatrix {
		if i == nBlocks-1 {
			backPaddingSize := int(b[blockLen-1])
			if backPaddingSize+paddingSize > len(b) {
				return "", errors.New("fail in depadding")
			} else {
				res = append(res, b[paddingSize:len(b)-backPaddingSize]...)
			}
		} else {
			res = append(res, b[paddingSize:]...)
		}
	}
	resStr := string(res)
	return resStr, nil
}
