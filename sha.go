package cryptor

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type sha struct {
	from   []byte
	length int32
}

func newSha(from []byte, length int32) *sha {
	return &sha{
		from:   from,
		length: length,
	}
}

func (s *sha) Hex() (to string) {
	bytes := s.Bytes()
	to = hex.EncodeToString(bytes[:])

	return
}

func (s *sha) Bytes() (to []byte) {
	switch s.length {
	case 256:
		bytes := sha256.Sum256(s.from)
		to = bytes[:]
	case 224:
		bytes := sha256.Sum224(s.from)
		to = bytes[:]
	case 384:
		bytes := sha512.Sum384(s.from)
		to = bytes[:]
	case 512:
		bytes := sha512.Sum512(s.from)
		to = bytes[:]
	}

	return
}
