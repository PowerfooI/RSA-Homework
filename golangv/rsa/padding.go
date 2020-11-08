package rsa

func PKCS1Padding(rawTextInput string, nBitSize int) *Blocks {
	nByteSize := nBitSize / 8
	paddingSize := 11
	pieceSize := nByteSize - paddingSize
	paddedBlocks := &Blocks{}

	input := []byte(rawTextInput)

	paddingContent := []byte{0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b}

	byteMatrix := make([][]byte, 0)
	lenOfInput := len(input)

	tmpSize := pieceSize
	for cur:=0;cur<lenOfInput; {
		if cur + pieceSize > lenOfInput {
				tmpSize = lenOfInput - cur
				newPaddingSize := nByteSize - tmpSize
				var newPaddingContent []byte
				for i:=0;i<newPaddingSize;i++ {
					newPaddingContent = append(newPaddingContent, byte(newPaddingSize))
				}
				tmpByteArray := make([]byte, tmpSize)
				copy(tmpByteArray, input[cur:cur+tmpSize])
				byteMatrix = append(byteMatrix, append(tmpByteArray, newPaddingContent...))
		} else {
			tmpByteArray := make([]byte, tmpSize)
			copy(tmpByteArray, input[cur:cur+tmpSize])
			byteMatrix = append(byteMatrix, append(tmpByteArray, paddingContent...))
		}
		cur += tmpSize
	}
	paddedBlocks.byteMatrix = byteMatrix
	return paddedBlocks
}

func PKCS1Depadding(blocks *Blocks, nBitSize int) string {
	//paddingSize := 11
	blockLen := len(blocks.byteMatrix[0])
	res := ""
	for _, b := range blocks.byteMatrix {
		paddingSize := int(b[blockLen-1])
		res += string(b[:len(b)-paddingSize])
	}
	return res
}
