package block

import (
	"crypto/cipher"
	"io"
)

// CopyBlockMode is like io.Copy but uses the given cipher.BlockMode to crypt each block before copying it to out.
func CopyBlockMode(blockMode cipher.BlockMode, out io.Writer, in io.Reader) error {
	// buffer for reading and crypting blocks
	buffer := make([]byte, blockMode.BlockSize())

	for {
		n, err := in.Read(buffer)
		if n == 0 && err == io.EOF {
			break
		}

		blockMode.CryptBlocks(buffer, buffer)

		_, err = out.Write(buffer)
		if err != nil {
			return err
		}
	}

	return nil
}
