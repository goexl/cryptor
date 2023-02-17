package cryptor

import (
	// nolint: gosec
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
)

type hmacBuilder[K typ] struct {
	hashFun hashFun
	key     K
	sign    []byte
}

func newHmacBuilder[K typ](key K, sign []byte) *hmacBuilder[K] {
	return &hmacBuilder[K]{
		hashFun: sha256.New,
		key:     key,
		sign:    sign,
	}
}

func (hb *hmacBuilder[K]) Md5() *hmacBuilder[K] {
	hb.hashFun = md5.New

	return hb
}

func (hb *hmacBuilder[K]) Sha512() *hmacBuilder[K] {
	hb.hashFun = sha512.New

	return hb
}

func (hb *hmacBuilder[K]) Build() *_hmac[K] {
	return newHmac[K](hb.hashFun, hb.key, hb.sign)
}
