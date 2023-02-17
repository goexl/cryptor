package cryptor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type _hmac[K typ] struct {
	hashFun hashFun
	key     K
	sign    []byte
}

func newHmac[K typ](hashFun hashFun, key K, sign []byte) *_hmac[K] {
	return &_hmac[K]{
		hashFun: hashFun,
		key:     key,
		sign:    sign,
	}
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
