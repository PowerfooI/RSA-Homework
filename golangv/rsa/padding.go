package rsa

import "math/rand"

var (
	paddingSize = 11
)

// 公钥加密类型的padding
func PKCS1Padding(rawTextInput string, nBitSize int) *Blocks {
	nByteSize := nBitSize / 8
	pieceSize := nByteSize - paddingSize
	paddedBlocks := &Blocks{}

	input := []byte(rawTextInput)

	byteMatrix := make([][]byte, 0)
	lenOfInput := len(input)

	tmpSize := pieceSize
	for cur := 0; cur < lenOfInput; {
		paddingContent := []byte{0x00, 0x02}
		// 11 - 2 - 1 = 8
		for i := 0; i < 8; i++ {
			paddingContent = append(paddingContent, byte(rand.Intn(0xff)))
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

func PKCS1Depadding(blocks *Blocks) string {
	nBlocks := len(blocks.byteMatrix)
	blockLen := len(blocks.byteMatrix[0])
	res := ""
	for i, b := range blocks.byteMatrix {
		if i == nBlocks-1 {
			backPaddingSize := int(b[blockLen-1])
			res += string(b[paddingSize : len(b)-backPaddingSize])
		} else {
			res += string(b[paddingSize:])
		}
	}
	return res
}
