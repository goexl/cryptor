package cryptor

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type _hmac[K typ] struct {
	hash func() hash.Hash
	key  K
	sign []byte
}

func newHmac[K typ](key K, sign []byte) *_hmac[K] {
	return &_hmac[K]{
		hash: sha256.New,
		key:  key,
		sign: sign,
	}
}

func (h *_hmac[K]) Md5() *_hmac[K] {
	h.hash = md5.New

	return h
}

func (h *_hmac[K]) String() string {
	return string(h.Sum())
}

func (h *_hmac[K]) Hex() string {
	return hex.EncodeToString(h.Sum())
}

func (h *_hmac[K]) Sum() (to []byte) {
	hashed := hmac.New(sha256.New, h.keyBytes())
	hashed.Write(h.sign)
	to = hashed.Sum(nil)

	return
}

func (h *_hmac[K]) keyBytes() (key []byte) {
	switch from := any(h.key).(type) {
	case string:
		key = []byte(from)
	case []byte:
		key = from
	}

	return
}
