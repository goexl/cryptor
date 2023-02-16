package cryptor

import (
	"crypto/sha256"
	"encoding/hex"
)

type _sha256 struct {
	from []byte
	typ  sha256Type
}

func newSha256(from []byte) *_sha256 {
	return &_sha256{
		from: from,
		typ:  sha256Type256,
	}
}

func (s *_sha256) Hex() (to string) {
	bytes := s.Bytes()
	to = hex.EncodeToString(bytes[:])

	return
}

func (s *_sha256) Bytes() (to []byte) {
	switch s.typ {
	case sha256Type256:
		bytes := sha256.Sum256(s.from)
		to = bytes[:]
	case sha256Type224:
		bytes := sha256.Sum224(s.from)
		to = bytes[:]
	}

	return
}

func (s *_sha256) For224() *_sha256 {
	s.typ = sha256Type224

	return s
}
