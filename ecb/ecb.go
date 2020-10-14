package ecb

import "crypto/cipher"

type ecb struct {
	block     cipher.Block
	blockSize int
}

// NewECBEncrypter returns a new cipher.BlockMode which uses the given Block cipher to encrypt given blocks in
// electronic code book mode (ECB). ECB is the most simple (and also most unsecure) mode in which a Block cipher
// can be operated. It simply encodes each block separately and totally lacks diffusion, which means that a single
// plaintext block is always encrypted to the exact same ciphertext block.
func NewECBEncrypter(block cipher.Block) cipher.BlockMode {
	return ecbEncrypter{block: block, blockSize: block.BlockSize()}
}

type ecbEncrypter ecb

func (e ecbEncrypter) BlockSize() int {
	return e.blockSize
}

func (e ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("grypto/ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("grypto/ecb: output smaller than input")
	}

	for len(src) > 0 {
		// encrypt in place with block cipher
		e.block.Encrypt(dst[:e.blockSize], src[:e.blockSize])

		// move to the next block
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}

// NewECBDecrypter returns a new cipher.BlockMode which uses the given Block cipher to decrypt given blocks in
// electronic code book mode (ECB). ECB is the most simple (and also most unsecure) mode in which a Block cipher
// can be operated. It simply encodes each block separately and totally lacks diffusion, which means that a single
// plaintext block is always encrypted to the exact same ciphertext block.
func NewECBDecrypter(block cipher.Block) cipher.BlockMode {
	return ecbDecrypter{block: block, blockSize: block.BlockSize()}
}

type ecbDecrypter ecb

func (e ecbDecrypter) BlockSize() int {
	return e.blockSize
}

func (e ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("grypto/ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("grypto/ecb: output smaller than input")
	}

	for len(src) > 0 {
		// encrypt in place with block cipher
		e.block.Decrypt(dst[:e.blockSize], src[:e.blockSize])

		// move to the next block
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}
