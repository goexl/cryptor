package cryptor

import (
	"crypto/md5"
	"encoding/hex"
)

type _md5 struct {
	from []byte
}

func newMd5(from []byte) *_md5 {
	return &_md5{
		from: from,
	}
}

func (m *_md5) Hex() (to string) {
	if bytes := m.Bytes(); nil != bytes {
		to = hex.EncodeToString(bytes)
	}

	return
}

func (m *_md5) Bytes() (to []byte) {
	hash := md5.New()
	if _, err := hash.Write(m.from); nil == err {
		to = hash.Sum(nil)
	}

	return
}
