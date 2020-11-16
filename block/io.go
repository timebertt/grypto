package block

import (
  "crypto/cipher"
  "errors"
  "io"
)

type reader struct {
  blockSize int
  blockMode cipher.BlockMode
  in        io.Reader
}

// NewBlockModeReader returns a io.Reader that crypts each block read from in with the given cipher.BlockMode.
func NewBlockModeReader(blockMode cipher.BlockMode, in io.Reader) io.Reader {
  return &reader{
    blockSize: blockMode.BlockSize(),
    blockMode: blockMode,
    in:        in,
  }
}

// Read implements io.Reader.
func (b *reader) Read(p []byte) (int, error) {
  // we must be able to read at least a full block
  if len(p) < b.blockSize {
    return 0, io.ErrShortBuffer
  }

  // read one block from in
  n, err := io.ReadFull(b.in, p[:b.blockSize])
  if err != nil {
    if err != io.EOF {
      return 0, err
    }
    // only return EOF, if we didn't read anything
    if n == 0 {
      return 0, err
    }
  }
  if n < b.blockSize {
    return 0, io.ErrUnexpectedEOF
  }

  func() {
    // cipher.BlockMode doesn't allow to return an error, so implementations tend to panic on error,
    // catch those errors here.
    defer func() {
      if p := recover(); p != nil {
        if e, ok := p.(error); ok {
          err = e
        } else if s, ok := p.(string); ok {
          err = errors.New(s)
        }
      }
    }()
    b.blockMode.CryptBlocks(p[:b.blockSize], p[:b.blockSize])
  }()
  if err != nil {
    return 0, err
  }
  return b.blockSize, nil
}
